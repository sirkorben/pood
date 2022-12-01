package helpers

import (
	"net/http"
	"net/mail"
	"pood/models"
	"regexp"
	"strings"
)

func ValidateUserData(w http.ResponseWriter, user *models.User) bool {
	space := regexp.MustCompile(`\s+`)
	user.FirstName = space.ReplaceAllString(strings.TrimSpace(user.FirstName), " ")
	user.LastName = space.ReplaceAllString(strings.TrimSpace(user.LastName), " ")
	user.Email = space.ReplaceAllString(strings.TrimSpace(user.Email), "")

	if user.FirstName == "" {
		ErrorResponse(w, FirstNameMissingErrorMsg, http.StatusBadRequest)
		return false
	}

	if user.LastName == "" {
		ErrorResponse(w, LastNameMissingErrorMsg, http.StatusBadRequest)
		return false
	}

	if user.Email == "" {
		ErrorResponse(w, EmailMissingErrorMsg, http.StatusBadRequest)
		return false
	}

	if user.Password == "" {
		ErrorResponse(w, PasswordMissingErrorMsg, http.StatusBadRequest)
		return false
	}

	if len(user.Password) < 6 || len(user.Password) > 20 {
		ErrorResponse(w, PasswordTooShortErrorMsg, http.StatusBadRequest)
		return false
	}

	_, errMail := mail.ParseAddress(user.Email)
	if errMail != nil {
		ErrorResponse(w, EmailIsNotValidErrorMsg, http.StatusBadRequest)
		return false
	}
	return true
}
