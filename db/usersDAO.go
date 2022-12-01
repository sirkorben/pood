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

	activated, err := activated(id)
	if err != nil {
		return 0, err
	}
	if activated != 1 {
		//return error that user is not activated   401
		return 0, models.ErrUserNotActivated
	}
	return id, nil
}

func activated(id int) (int, error) {
	var activated int
	row := DB.QueryRow("select activated from users where id = ?", id)
	err := row.Scan(&activated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 500, models.ErrNoRecord
		} else {
			return 500, err
		}
	}
	return activated, nil
}

func GetFirstNameById(id int) (*models.User, error) {
	row := DB.QueryRow("select id, firstname from users where id = ?", id)
	u := &models.User{}
	err := row.Scan(&u.Id, &u.FirstName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return u, nil
}
