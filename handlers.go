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
	if r.URL.Path != "/" {
		helpers.ErrorResponse(w, helpers.NotFoundErrorMsg, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodOptions {
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
			var errMsg *helpers.ErrorMsg
			if errors.As(err, &errMsg) {
				helpers.ErrorResponse(w, *errMsg, http.StatusBadRequest)
			} else {
				log.Println(err)
				helpers.ErrorResponse(w, *errMsg, http.StatusInternalServerError)
			}
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
	if r.Method == http.MethodPost {
		var u models.User
		err := helpers.DecodeJSONBody(w, r, &u)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			// var errMsg *helpers.ErrorMsg
			// if errors.As(err, &errMsg) {
			// 	helpers.ErrorResponse(w, *errMsg, http.StatusBadRequest)
			// } else {
			// 	log.Println("helpers.DecodeJSONBody(w, r, &u)", err)
			// 	helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			// }
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
				log.Println("db.Authenticate(credential, u.Password)", err)
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
			log.Println(err.Error())
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
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
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

func search(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		helpers.ErrorResponse(w, helpers.MethodNotAllowedErrorMsg, http.StatusMethodNotAllowed)
		return
	}

	_, err := db.CheckSession(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodPost {
		// "article": "045121011hx"  "article": "4M2820160"
		var article models.Article
		err := helpers.DecodeJSONBody(w, r, &article)
		if err != nil {
			helpers.HandleDecodeJSONBodyError(err, w)
			return
		}
		// increse the price by 40% - it will be choosen different discount taken from User info set by admin
		// would be taken from here knowing whos is logged take his percent from field(TODO: add field to users)
		percent := 1.4
		var prices models.ApiResponse
		err = helpers.CallForPrices(article.Article, &prices)
		if err != nil {
			// handle error
			log.Println(err)
			return
		}

		for _, obj := range prices.Prices {
			models.ChangePrice(obj, percent)
		}
		helpers.WriteResponse(prices, w) // check for possible errors
	}

}

// admin logic
func admin(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		helpers.ErrorResponse(w, helpers.MethodNotAllowedErrorMsg, http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodGet || r.Method == http.MethodPatch {
		_, err := db.CheckAdminSession(r)
		if err != nil {
			helpers.ErrorHandler(err, w)
			return
		}
	} else {
		helpers.ErrorResponse(w, helpers.MethodNotAllowedErrorMsg, http.StatusMethodNotAllowed)
	}
}

func adminApproveHandler(w http.ResponseWriter, r *http.Request) {
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

}

func userOrders(w http.ResponseWriter, r *http.Request) {

}

func addProductToCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	user := &models.User{
		Id: -1,
	}
	s, err := db.CheckSession(r)
	if err != nil {
		// handle better
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	} else {
		user = s.User
		log.Println(fmt.Sprintf("User with [Id - %v] accessed to POST /myorders/add", s.User.Id), user.FirstName)
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

func shoppingCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	user := &models.User{
		Id: -1,
	}
	s, err := db.CheckSession(r)
	if err != nil {
		// handle better
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	} else {
		// could be deleted ??
		user = s.User
		log.Println(fmt.Sprintf("User with [Id - %v] accessed to POST /cart", s.User.Id), user.FirstName)
	}
	if r.Method == http.MethodGet {
		// show existing non confirmed order
		var shoppingCart models.ShoppingCart
		shoppingCart.Products, err = db.GetProductsUnderNonConfirmedOrderId(s.User.Id)
		if err != nil {
			fmt.Println("\t1")
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
		}
		models.SumPrices(&shoppingCart)
		helpers.WriteResponse(shoppingCart, w) // check for possible errors
	}
	if r.Method == http.MethodPost {
		// confirm existing order
		// create new non confirmed order for future
		err = db.ConfirmOrderId(s.User.Id)
		if err != nil {
			helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
		}
	}
}
