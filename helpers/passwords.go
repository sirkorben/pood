package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		// handle error
		log.Println(err)
		return "", err
	}
	return string(hashedPassword), nil
}
