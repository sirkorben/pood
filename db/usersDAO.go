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

	_, err = DB.Exec("INSERT INTO users (firstname, lastname, email, password) VALUES (?,?,?,?)",
		firstName, lastName, email, hashedPassword)
	if err != nil {
		log.Println("sqlite.InsertUser()", err)
		return err
	}
	return nil
}
