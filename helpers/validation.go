//TODO: Handle errors !!!!!!!!!!!!S
package helpers

import (
	"net/http"
	"net/mail"
	"pood/models"
	"regexp"
	"strings"
)

func ValidateUserData(w http.ResponseWriter, user *models.User) bool {
	var errMsg ErrorMsg
	space := regexp.MustCompile(`\s+`)
	user.FirstName = space.ReplaceAllString(strings.TrimSpace(user.FirstName), " ")
	user.LastName = space.ReplaceAllString(strings.TrimSpace(user.LastName), " ")
	user.Email = space.ReplaceAllString(strings.TrimSpace(user.Email), "")

	if user.FirstName == "" {
		errMsg.ErrorDescription = "Firstname is missing."
		errMsg.ErrorType = "FIRSTNAME_FIELD_EMPTY"
		ErrorResponse(w, http.StatusBadRequest)
		return false
	}

	if user.LastName == "" {
		errMsg.ErrorDescription = "Lastname is missing."
		errMsg.ErrorType = "LASTNAME_FIELD_EMPTY"
		ErrorResponse(w, http.StatusBadRequest)
		return false
	}

	if user.Email == "" {
		errMsg.ErrorDescription = "Email is missing."
		errMsg.ErrorType = "EMAIL_FIELD_EMPTY"
		ErrorResponse(w, http.StatusBadRequest)
		return false
	}

	if user.Password == "" {
		errMsg.ErrorDescription = "Password is missing."
		errMsg.ErrorType = "PASSWORD_FIELD_EMPTY"
		ErrorResponse(w, http.StatusBadRequest)
		return false
	}

	if len(user.Password) < 6 || len(user.Password) > 20 {
		errMsg.ErrorDescription = "Password is too short - 6 chars min."
		errMsg.ErrorType = "PASSWORD_TOO_SHORT"
		ErrorResponse(w, http.StatusBadRequest)
		return false
	}

	_, errMail := mail.ParseAddress(user.Email)
	if errMail != nil {
		errMsg.ErrorDescription = "Email is not valid"
		errMsg.ErrorType = "EMAIL_INVALID"
		ErrorResponse(w, http.StatusBadRequest)
		return false
	}
	return true
}
