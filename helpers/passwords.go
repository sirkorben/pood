package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		// handle error
	}
	return string(hashedPassword)
}
