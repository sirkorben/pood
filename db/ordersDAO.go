package db

import (
	"database/sql"
	"errors"
	"log"

	uuid "github.com/satori/go.uuid"
)

func CreateNonConfirmedOrderByUserEmail(email string) error {
	orderId := uuid.NewV4()
	_, err := DB.Exec("INSERT INTO orders (id, user_id, confirmed, date_created) VALUES (?,(SELECT id FROM users WHERE email = ?),?,strftime('%s','now'))",
		orderId, email, 0)
	if err != nil {
		log.Println("sqlite.orders err \t", err)
		return err
	}
	return nil
}

func GetNonConfirmedOrderId(userId int) (string, error) {
	var orderId string
	row := DB.QueryRow("SELECT id FROM orders WHERE user_id = ? AND confirmed = 0", userId)
	err := row.Scan(&orderId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("1", err)
			// handle error
			// should not happen ???
			return "", err
		} else {
			log.Println("2", err)
			return "", err
		}
	}
	return orderId, nil
}

func GetConfirmedOrderIds(userId int) ([]string, error) {
	rows, err := DB.Query("select id, user_id from orders WHERE confirmed = 1 ORDER BY date_created desc")
	if err != nil {
		// handle error
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var orderIds []string
	for rows.Next() {
		var orderId string
		err = rows.Scan(&orderId)
		if err != nil {
			// handle err
			// return nil, err
			log.Println(err)
			return nil, err
		}
		orderIds = append(orderIds, orderId)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return orderIds, nil
}

func GetConfirmedOrderIdsForAdmin() ([]string, error) {
	rows, err := DB.Query("select id, user_id from orders WHERE confirmed = 1 ORDER BY date_created desc")
	if err != nil {
		// handle error
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var orderIds []string
	for rows.Next() {
		var orderId string
		err = rows.Scan(&orderId)
		if err != nil {
			// handle err
			// return nil, err
			log.Println(err)
			return nil, err
		}
		orderIds = append(orderIds, orderId)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return orderIds, nil
}

func ConfirmOrderId(userId int) error {

	return nil
}
