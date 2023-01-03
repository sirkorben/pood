package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"pood/helpers"
	"pood/models"

	uuid "github.com/satori/go.uuid"
)

const default_user_percent = 1.15

func InsertUser(user models.UserRegistered) error {
	firstName := user.FirstName
	lastName := user.LastName
	email := user.Email
	password := user.Password
	hashedPassword, err := helpers.GeneratePassword(password)
	if err != nil {
		return err
	}

	row := DB.QueryRow("SELECT id FROM users WHERE email = ?", email)
	err = row.Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		return models.ErrDuplicateEmail
	}
	// insert user to db
	_, err = DB.Exec("INSERT INTO users (firstname, lastname, email, password, is_admin, activated, user_percent, date_created) VALUES (?,?,?,?,?,?,?, strftime('%s','now'))",
		firstName, lastName, email, hashedPassword, 0, 0, default_user_percent)
	if err != nil {
		log.Println("sqlite.InsertUser()", err)
		return err
	}
	// insert his order number and confirmed = 0. when it will be confirmed then create a new one
	// can be separate function where checks are done in ordersDAO
	orderId := uuid.NewV4()
	_, err = DB.Exec("INSERT INTO orders (id, user_id, confirmed, date_created) VALUES (?,(SELECT id FROM users WHERE email = ?),?,strftime('%s','now'))",
		orderId, email, 0)
	if err != nil {
		log.Println("sqlite.orders err \t", err)
		return err
	}
	return nil
}

func Authenticate(credName, password string) (int, error) {
	var id int
	var hashedPassword []byte
	row := DB.QueryRow("SELECT id, password FROM users WHERE email = ?", credName)
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

	isActivated, err := isActivated(id)
	if err != nil {
		return 0, err
	}
	if isActivated != 1 {
		return 0, models.ErrUserNotActivated
	}
	return id, nil
}

func isActivated(id int) (int, error) {
	var activated int
	row := DB.QueryRow("SELECT activated FROM users WHERE id = ?", id)
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
	row := DB.QueryRow("SELECT id, firstname FROM users WHERE id = ?", id)
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

func GetNonActivatedUsers() ([]*models.User, error) {
	var users []*models.User

	rows, err := DB.Query("SELECT id, firstname, lastname, email, activated, date_created FROM users WHERE activated = 0 ORDER BY date_created DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Activated, &user.DateCreated)
		if err != nil {

			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetActivatedUsers() ([]*models.User, error) {
	var users []*models.User

	rows, err := DB.Query("SELECT id, firstname, lastname, email, activated, date_created, user_percent FROM users WHERE activated = 1 ORDER BY date_created DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Activated, &user.DateCreated, &user.UserPercent)
		if err != nil {

			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func ActivateUser(id int) error {
	_, err := DB.Exec("UPDATE users SET activated = 1 WHERE id = ?;", id)
	if err != nil {
		fmt.Println("4	", err)
		return err
	}
	return nil
}

func GetPercentByUserId(id int) float64 {
	row := DB.QueryRow("SELECT user_percent FROM users WHERE id = ?", id)
	u := &models.User{}
	err := row.Scan(&u.UserPercent)
	if err != nil {
		return 1.15
	}
	return *u.UserPercent
}

func ManageUserPercent(id int, percent float64) error {
	_, err := DB.Exec("UPDATE users SET user_percent = ? WHERE id = ?;", percent, id)
	if err != nil {
		fmt.Println("5	", err)
		return err
	}
	return nil
}
