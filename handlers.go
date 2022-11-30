package main

import (
	"errors"
	"log"
	"net/http"
	"pood/db"
	"pood/helpers"
	"pood/models"
)

//"http://localhost:3000"
func enableCors(w *http.ResponseWriter) {
	//(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500") // changed for my live server addr
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // changed for my live server addr
	
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Authorization, Accept")
}

func home(w http.ResponseWriter, r *http.Request) {
	// will be used to serve opportunity to signin or signup
	enableCors(&w)
	if r.URL.Path != "/" {
		helpers.ErrorResponse(w, http.StatusBadRequest)
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
				helpers.ErrorResponse(w, http.StatusBadRequest)
			} else {
				log.Println(err)
				helpers.ErrorResponse(w, http.StatusInternalServerError)
			}
			return
		}
		if helpers.ValidateUserData(w, &u) {
			err := db.InsertUser(u)
			if err != nil {
				var errMsg *helpers.ErrorMsg
				if errors.Is(err, models.ErrDuplicateUsername) {
					errMsg.ErrorDescription = "Username already taken."
					errMsg.ErrorType = "USERNAME_ALREADY_TAKEN"
					helpers.ErrorResponse(w, http.StatusBadRequest)
					return
				}

				if errors.Is(err, models.ErrDuplicateEmail) {
					errMsg.ErrorDescription = "Email already taken."
					errMsg.ErrorType = "EMAIL_ALREADY_TAKEN"
					helpers.ErrorResponse(w, http.StatusBadRequest)
					return
				}

				helpers.ErrorResponse(w, http.StatusInternalServerError)
				return
			} else {
				log.Println("User inserted - ", u.Email)
			}
		}
	}
}

func signIn(w http.ResponseWriter, r *http.Request) {

}

func search(w http.ResponseWriter, r *http.Request) {

}
