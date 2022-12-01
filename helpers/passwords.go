package helpers

import (
	"errors"
	"log"
	"pood/models"

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

func CheckPassword(hashedPass []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.ErrInvalidCredentials
		} else {
			//return InternalServerError? from models var ??
			return err
		}
	}
	return nil
}
