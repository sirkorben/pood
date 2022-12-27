package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"pood/models"

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
			log.Println("GetNonConfirmedOrderId err 1 ->", err)
			return "", err
		} else {
			log.Println("GetNonConfirmedOrderId err 2 ->", err)
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

func ConfirmOrder(userId int) (string, error) {
	orderId, err := GetNonConfirmedOrderId(userId)
	if err != nil {
		// handle err
		log.Println("ConfirmOrder err 1", err)
		return "", err
	}

	sqlStatement := `SELECT position_id FROM positions_ordered WHERE order_id=$1;`
	var positionId string

	row := DB.QueryRow(sqlStatement, orderId)
	err = row.Scan(&positionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", models.ErrNoRecord
		}
	} else {
		err2 := updateOrderConfirmedField(orderId)
		if err2 != nil {
			return "", err2
		}
		err2 = createNewNonConfirmedOrderByUserId(userId)
		if err2 != nil {
			return "", err2
		}
	}
	log.Printf("[order_id: %v] created", orderId)
	return orderId, nil
}

func updateOrderConfirmedField(orderId string) error {
	_, err := DB.Exec("UPDATE orders SET confirmed = 1, date_created = strftime('%s','now') WHERE id = ?;", orderId)
	if err != nil {
		fmt.Println("updateOrderConfirmedField err ->", err)
		return err
	}
	return nil
}

func createNewNonConfirmedOrderByUserId(userId int) error {
	orderId := uuid.NewV4()
	_, err := DB.Exec("INSERT INTO orders (id, user_id, confirmed, date_created) VALUES (?,?,?,strftime('%s','now'))",
		orderId, userId, 0)
	if err != nil {
		log.Println("createNonConfirmedOrderByUserId err \t", err)
		return err
	}
	return nil
}

func DeleteShoppingCart(orderId string) error {
	_, err := DB.Exec("delete from positions_ordered where order_id = ?", orderId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("DeleteShoppingCartByOrderId err -> ", err)
		return err
	}
	return nil
}

func DeletePositionFromCart(positionId string) error {
	_, err := DB.Exec("delete from positions_ordered where position_id = ?", positionId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("DeleteShoppingCartByOrderId err -> ", err)
		return err
	}
	return nil
}

// func GetConfirmedUserOrders(userId int) ([]string, error) {
// 	rows, err := DB.Query("select id from orders WHERE confirmed = 1 ORDER BY date_created desc")
// 	if err != nil {
// 		// handle error
// 		log.Println(err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var userOrders []string
// 	for rows.Next() {
// 		var orderId string
// 		err = rows.Scan(&orderId)
// 		if err != nil {
// 			// handle err
// 			// return nil, err
// 			log.Println(err)
// 			return nil, err
// 		}
// 		userOrders = append(userOrders, orderId)
// 	}
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return userOrders, nil
// }

func GetConfirmedUserOrders2(userId int) ([]*models.UserOrder, error) {
	rows, err := DB.Query("select id, date_created from orders WHERE confirmed = 1 ORDER BY date_created desc")
	if err != nil {
		// handle error
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	// var userOrders []string
	var userOrders []*models.UserOrder
	for rows.Next() {
		order := &models.UserOrder{}
		err = rows.Scan(&order.OrderId, &order.DateCreated)
		if err != nil {
			// handle err
			// return nil, err
			log.Println(err)
			return nil, err
		}
		userOrders = append(userOrders, order)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return userOrders, nil
}

func GetConfirmedUserOrder(orderId string) (models.UserOrder, error) {
	var order models.UserOrder

	return order, nil
}

func GetOrderDateCreated(userId int, orderId string) (int, error) {
	var dateCreated int
	row := DB.QueryRow("SELECT date_created FROM orders WHERE id = ? AND user_id = ?", orderId, userId)
	err := row.Scan(&dateCreated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("GetNonConfirmedOrderId err 1 ->", err)
			return 0, err
		} else {
			log.Println("GetNonConfirmedOrderId err 2 ->", err)
			return 0, err
		}
	}
	return dateCreated, nil
}
