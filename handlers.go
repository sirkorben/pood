package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pood/db"
	"pood/helpers"
	"pood/models"
	"strings"

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
			var errMsg *helpers.ErrorMsg
			if errors.As(err, &errMsg) {
				helpers.ErrorResponse(w, *errMsg, http.StatusBadRequest)
			} else {
				log.Println("helpers.DecodeJSONBody(w, r, &u)", err)
				helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
			}
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
		return
	}
	user := &models.User{
		Id: -1,
	}
	s, err := db.CheckSession(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	} else {
		user = s.User
		log.Println(fmt.Sprintf("User with [Id - %v] accessed to GET /search", s.User.Id), user.FirstName)
	}

	// No need to call API everytime

	// if r.Method == http.MethodPost {
	// 	// TODO: provide CallForPrices with article given by client POST /search
	// 	article := "045121011hx" // delete it
	// 	// increse the price by 40% - it will be choosen different discount taken from User info set by admin
	// 	discount := 1.4
	// 	var prices models.ApiResponse
	// 	err := helpers.CallForPrices(article, &prices)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}

	// 	for _, obj := range prices.Prices {
	// 		models.ChangePrice(obj, discount)
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	jsonResp, err := json.Marshal(prices)
	// 	if err != nil {
	// 		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// 		helpers.ErrorResponse(w, helpers.InternalServerErrorMsg, http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Write(jsonResp)
	// }

}

// admin logic
func adminHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}
	_, err := db.CheckAdminSession(r)
	if err != nil {
		if errors.Is(err, models.ErrUnauthorized) {
			helpers.ErrorResponse(w, helpers.StatusForbiddenErrorMsg, http.StatusForbidden)
		} else if errors.Is(err, models.ErrNoRecord) {
			helpers.ErrorResponse(w, helpers.BadRequestErrorMsg, http.StatusBadRequest)
		} else {
			log.Println(err)
			helpers.ErrorResponse(w, helpers.UserNotActivatedErrorMsg, http.StatusUnauthorized)
		}
		return
	} else {
		log.Println("Admin - entered GET /admin")
	}
	url := strings.Split(strings.Trim(r.URL.Path, "/"), "/") // 	/admin/approve/ -> [admin, approve]
	if len(url) == 2 && url[1] == "approve" {
		if r.Method == http.MethodGet {
			log.Println("Admin - listing non activated users on GET", r.URL.Path)
			nonActivatedUsers()
		}
		if r.Method == http.MethodPatch {
			log.Println("Admin - activating users on POST", r.URL.Path)
			adminActivateUser()
		}
	}
	if len(url) == 2 && url[1] == "orders" { //		/admin/orders/ -> [admin, orders]
		log.Println("Admin - listing orders", r.URL.Path)

	}
}

func nonActivatedUsers() ([]*models.User, error) {
	nonactivatedUsers, err := db.GetNonActivatedUsers()
	if err != nil {
		// handle error
		// handle situation of NoRows
		return nil, err
	}
	return nonactivatedUsers, err
}

func adminActivateUser() {
	log.Println("assume we have seen unactivated user, and sending request to update status of chosen user")
}
