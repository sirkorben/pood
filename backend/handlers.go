package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pood/db"
	"pood/helpers"
	"pood/models"

	uuid "github.com/satori/go.uuid"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Authorization, Accept")
}

func home(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	if r.URL.Path != "/" {
		helpers.ErrorResponse(w, helpers.NotFoundErrorMsg, http.StatusNotFound)
		return
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method == http.MethodPost {
		var u models.User
		err := helpers.DecodeJSONBody(w, r, &u)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			return
		}
		if helpers.ValidateUserData(w, &u) {
			err := db.InsertUser(u)
			if err != nil {
				if errors.Is(err, models.ErrDuplicateEmail) {
					helpers.ErrorResponse(w, helpers.EmailAlreadyTakenErrorMsg, http.StatusBadRequest)
					return
				}
				log.Println(err.Error())
				helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
				return
			} else {
				log.Println("User inserted - ", u.Email)
			}
		}
	}
}

func signIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}
	log.Println("somebody signed in after cors check")
	if r.Method == http.MethodPost {
		var u models.User
		err := helpers.DecodeJSONBody(w, r, &u)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			return
		}

		credential := u.Email
		id, err := db.Authenticate(credential, u.Password)
		if err != nil {
			if errors.Is(err, models.ErrInvalidCredentials) {
				helpers.ErrorResponse(w, helpers.CredentialsDontMatchErrorMsg, http.StatusBadRequest)
			} else if errors.Is(err, models.ErrUserNotActivated) {
				helpers.ErrorResponse(w, helpers.UserNotActivatedErrorMsg, http.StatusUnauthorized)
			} else {
				helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			}
			return
		}

		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:   "session",
			Value:  sID.String(),
			MaxAge: 60 * 60 * 24,
		}
		http.SetCookie(w, c)
		err = db.InsertSession(c.Value, id)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}

		log.Printf("User with [Id - %v] joined the Pood", id)
	}
}

func signOut(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}
	s, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		err = db.DeleteSession(s.Id)
		if err != nil {
			log.Println(err.Error())
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		})
	}

}

func search(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	user := &models.User{
		Id: -1,
	}
	s, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	} else {
		// could be deleted ?? // left here as an example if I need User
		user = s.User
		log.Println(fmt.Sprintf("User with [Id - %v] accessed /search endpoint", s.User.Id), user.FirstName)
	}

	if r.Method == http.MethodPost {
		// "article": "045121011hx"  "article": "4M2820160"
		var article models.Article
		err := helpers.DecodeJSONBody(w, r, &article)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			return
		}
		if helpers.ValidateSearchByArticle(article) {
			// TODO:
			// increse the price by 40% - it will be choosen different discount taken from User info set by admin
			// would be taken from here knowing whos is logged take his percent from field(TODO: add field to users)
			percent := 1.4
			var prices models.ApiResponse
			err = helpers.ApiCall(article.Article, &prices)
			if err != nil {
				log.Println(err)
				helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
				return
			}

			for _, obj := range prices.Prices {
				models.ChangePrice(obj, percent)
			}
			helpers.WriteResponse(prices, w) // check for possible errors
		} else {
			helpers.ErrorResponse(w, helpers.BadRequestErrorMsg, http.StatusBadRequest)
		}
	}
}

// admin logic
func admin(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	_, err := db.CheckAdminSession(r)
	if err != nil {
		helpers.ErrorHandler(err, w)
		return
	}
}

func adminApproveHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	_, err := db.CheckAdminSession(r)
	if err != nil {
		helpers.ErrorHandler(err, w)
		return
	}
	if r.Method == http.MethodGet {
		nonActivatedUsersList, err := db.GetNonActivatedUsers()
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				// handle error
				log.Println("adminApproveHandler -> MethodGet -> nonActivatedUserList, err := ")
				helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
				return
			}
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}
		err = helpers.WriteResponse(nonActivatedUsersList, w)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
		}

	} else if r.Method == http.MethodPatch {
		var userToActivate models.User
		err := helpers.DecodeJSONBody(w, r, &userToActivate)
		if err != nil {
			var errMsg *helpers.ErrorMsg
			if errors.As(err, &errMsg) {
				helpers.ErrorResponse(w, *errMsg, http.StatusBadRequest)
			} else {
				log.Println("helpers.DecodeJSONBody(w, r, &u)", err)
				helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			}
			return
		}
		err = db.ActivateUser(userToActivate)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
		}
	} else {
		helpers.ErrorResponse(w, helpers.MethodNotAllowedErrorMsg, http.StatusMethodNotAllowed)
	}
}

func adminOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// show all users' orders
}

func userOrders(w http.ResponseWriter, r *http.Request) {
	// show user confirmed orders
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	s, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		var userOrders models.UserOrders
		userOrders.Orders, err = db.GetConfirmedUserOrders2(s.User.Id)
		if err != nil {
			helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
			return
		}
		helpers.WriteResponse(userOrders, w)
	}
}

func order(w http.ResponseWriter, r *http.Request) {
	// show user confirmed order by query param /order?id=8d6a4012-98e9-4a38-82e3-c27f6fbbf419
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	s, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	orderId := r.URL.Query().Get("id") // handle execptions
	var order models.UserOrder
	order.Positions, err = db.GetProductsUnderNonConfirmedOrderId(s.User.Id, orderId)
	if err != nil {
		helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
		return
	}
	dateCreated, err := db.GetOrderDateCreated(s.User.Id, orderId)
	models.CollectUserOrder(&order, orderId, dateCreated)
	if err != nil {
		helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(order, w)
}

func shoppingCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	s, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		// show existing non confirmed order for user
		// TODO: add empty non confirmed order if doesnt exist
		var shoppingCart models.ShoppingCart
		orderId, err := db.GetNonConfirmedOrderId(s.User.Id)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}
		shoppingCart.Products, err = db.GetProductsUnderNonConfirmedOrderId(s.User.Id, orderId)
		if err != nil {
			fmt.Println("\t1")
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}
		models.CreateShoppingCart(&shoppingCart, orderId)
		helpers.WriteResponse(shoppingCart, w)
	}
}

func addItemToCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	s, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		var productToAddIntoCart models.Product
		err := helpers.DecodeJSONBody(w, r, &productToAddIntoCart)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			return
		}

		err = db.AddProductToShoppingCart(s.User.Id, productToAddIntoCart)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
		}
	}
}

func confirmCart(w http.ResponseWriter, r *http.Request) {

	//	TODO: send information about confirmed order to client and admin emails

	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	s, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		orderId, err := db.ConfirmOrder(s.User.Id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				helpers.ErrorResponse(w, helpers.EmptyCartErrorMsg, http.StatusBadRequest)
			} else {
				helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			}
			return
		}

		var order models.UserOrder
		order.Positions, err = db.GetProductsUnderNonConfirmedOrderId(s.User.Id, orderId)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}
		dateCreated, err := db.GetOrderDateCreated(s.User.Id, orderId)
		models.CollectUserOrder(&order, orderId, dateCreated)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}

		err = helpers.SendEmail(order)
		if err != nil {
			log.Println("helpers.SendEmail err -> ", err)
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			return
		}
		log.Println("email sent")
		helpers.InfoResponse(w, helpers.OrderConfirmedInfoMsg, http.StatusCreated)

	}
}

func removeCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	_, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodDelete {
		var shoppingCartOrderId models.OrderId
		err := helpers.DecodeJSONBody(w, r, &shoppingCartOrderId)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			return
		}
		db.DeleteShoppingCart(shoppingCartOrderId.OrderId)
	}
}

func removeItemFromCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}

	_, err := db.CheckSession(r)
	if err != nil {
		helpers.ErrorResponse(w, helpers.UnauthorizedErrorMsg, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodDelete {
		var shoppingCartPositionId models.PositionId
		err := helpers.DecodeJSONBody(w, r, &shoppingCartPositionId)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			return
		}
		db.DeletePositionFromCart(shoppingCartPositionId.PositionId)
	}
}
