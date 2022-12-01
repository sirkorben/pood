package main

import (
	"errors"
	"log"
	"net/http"
	"pood/db"
	"pood/helpers"
	"pood/models"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000") // changed for my live server addr
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Authorization, Accept")
}

func home(w http.ResponseWriter, r *http.Request) {
	// will be used to serve opportunity to signin or signup
	enableCors(&w)
	if r.URL.Path != "/" {
		helpers.ErrorResponse(w, helpers.NotFoundErrorMsg, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodOptions {
		return
	}
}

// after signup should be message that web site will be available after Admin verify user
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
				// var errMsg *helpers.ErrorMsg
				// if errors.Is(err, models.ErrDuplicateUsername) {
				// 	errMsg.ErrorDescription = "Username already taken."
				// 	errMsg.ErrorType = "USERNAME_ALREADY_TAKEN"
				// 	helpers.ErrorResponse(w, http.StatusBadRequest)
				// 	return
				// }
				if errors.Is(err, models.ErrDuplicateEmail) {
					helpers.ErrorResponse(w, helpers.EmailAlreadyTakenErrorMsg, http.StatusBadRequest)
					return
				}

				helpers.ErrorResponse(w, helpers.InternalErrorMsg, http.StatusInternalServerError)
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
	// if r.Method == http.MethodPost {
	// 	var u models.User
	// 	err := helpers.DecodeJSONBody(w, r, &u)
	// 	if err != nil {
	// 		var errMsg *helpers.ErrorMsg
	// 		if errors.As(err, &errMsg) {
	// 			errorResponse(w, *errMsg, http.StatusBadRequest)
	// 		} else {
	// 			log.Println(err.Error())
	// 			errorResponse(w, internalErrorMsg, http.StatusInternalServerError)
	// 		}
	// 		return
	// 	}
	// 	var credential string
	// 	if u.Email == "" {
	// 		credential = u.Username
	// 	} else {
	// 		credential = u.Email
	// 	}
	// 	id, err := sqlite.Authenticate(credential, u.Password)
	// 	if err != nil {
	// 		var errMsg ErrorMsg
	// 		if errors.Is(err, models.ErrInvalidCredentials) {
	// 			errMsg.ErrorDescription = "Email/username and password don't match."
	// 			errMsg.ErrorType = "CREDENTIALS_DONT_MATCH"
	// 			errorResponse(w, errMsg, http.StatusBadRequest)
	// 		} else {
	// 			log.Println(err.Error())
	// 			errorResponse(w, internalErrorMsg, http.StatusInternalServerError)
	// 		}
	// 		return
	// 	}

	// 	sID := uuid.NewV4()
	// 	c := &http.Cookie{
	// 		Name:   "session",
	// 		Value:  sID.String(),
	// 		MaxAge: 60 * 60 * 24,
	// 	}
	// 	http.SetCookie(w, c)

	// 	err = sqlite.InsertSession(c.Value, id)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		errorResponse(w, internalErrorMsg, http.StatusInternalServerError)
	// 		return
	// 	}
	// }
}

func search(w http.ResponseWriter, r *http.Request) {

}
