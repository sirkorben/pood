package db

import (
	"database/sql"
	"errors"
	"log"
	"pood/helpers"
	"pood/models"
)

func InsertUser(user models.User) error {
	firstName := user.FirstName
	lastName := user.LastName
	email := user.Email
	password := user.Password
	hashedPassword, err := helpers.GeneratePassword(password)
	if err != nil {
		return err
	}

	row := DB.QueryRow("select id from users where email = ?", email)
	err = row.Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		return models.ErrDuplicateEmail
	}

	_, err = DB.Exec("INSERT INTO users (firstname, lastname, email, password, is_admin, activated) VALUES (?,?,?,?,?,?)",
		firstName, lastName, email, hashedPassword, 0, 0)
	if err != nil {
		log.Println("sqlite.InsertUser()", err)
		return err
	}
	return nil
}

func Authenticate(credName, password string) (int, error) {
	var id int
	var hashedPassword []byte
	row := DB.QueryRow("select id, password from users where email = ?", credName)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = helpers.CheckPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	return id, nil
}
