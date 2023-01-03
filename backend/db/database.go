//TODO 1: Hide Admin Password to Local Env or find the way it should be
//TODO 2: setup proper logging
//TODO 3: handle errors properly with logging regarding db errors
//TODO later: think about possibility to use real db if it is needed ?

//HINTS: probably while hosting app we need to set up environment beforeand go get github.com/mattn/go-sqlite3

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pood/helpers"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

var DB = &sql.DB{}

const dbName = "pood.db"

func InitDatabase() {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
		return
	}

	DB = db
	if !fileExists(fmt.Sprintf("./%v", dbName)) {
		fillDbWithTablesAndAdmin()
	}

	if err = DB.Ping(); err != nil {
		// handle error
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fillDbWithTablesAndAdmin() {
	adminEmail := os.Getenv("POOD_ADMIN_EMAIL")
	adminPass, err := helpers.GeneratePassword(os.Getenv("POOD_ADMIN_PASSWORD"))
	if err != nil {
		// handle error
		return
	}
	// may not product fields be not null ????
	const CREATE string = `
	create table users(	
		id INTEGER not null primary key autoincrement,
		firstname TEXT not null,
		lastname TEXT not null,
		email TEXT not null unique, 
		password BLOB not null,
		is_admin INTEGER not null,
		activated INTEGER not null,
		user_percent REAL not null,
		date_created INTEGER not null);

	create table orders(
		id TEXT not null primary key unique, 
		user_id INTEGER not null,
		confirmed INTEGER not null, 
		date_created INTEGER not null);

	create table positions_ordered(
		position_id INTEGER not null primary key autoincrement,
		order_id TEXT not null,
		price REAL not null,
		article TEXT not null,
		supplier TEXT not null,
		supplier_price_num TEXT not null,
		brand TEXT not null,
		currency TEXT not null,
		currency_rate TEXT not null,
		delivery TEXT not null,
		weight REAL not null,
		quantity INTEGER not null,
		product_quantity_price REAL not null);	

	create table sessions(
		id TEXT not null primary key, 
		user_id INTEGER not null unique, 
		date_created INTEGER not null);
	`

	_, err = DB.Exec(CREATE)
	if err != nil {
		// handle error
		// error comes here while starting application with created database is in root
		log.Println("error in _, err = DB.Exec(CREATE)", err)
		return
	}

	_, err = DB.Exec("INSERT INTO users (id, firstname, lastname, email, password, is_admin, activated, user_percent, date_created) VALUES (?,?,?,?,?,?,?,?, strftime('%s','now'))",
		1, "Daniil", "Batjkovich", adminEmail, adminPass, 1, 1, 1.0)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(INSERT) 1\n ", err)
		return
	}

	orderIdForAdmin := uuid.NewV4()
	_, err = DB.Exec("INSERT INTO orders (id, user_id, confirmed, date_created) VALUES (?,(SELECT id FROM users WHERE email = ?),?,strftime('%s','now'))",
		orderIdForAdmin, adminEmail, 0)
	if err != nil {
		log.Println("sqlite.orders err \t", err)
	}

	// delete users below
	oneTwoThree, err := helpers.GeneratePassword("123456")
	if err != nil {
		// handle error
		return
	}
	_, err = DB.Exec("INSERT INTO users (id, firstname, lastname, email, password, is_admin, activated, user_percent, date_created) VALUES (?,?,?,?,?,?,?,?, strftime('%s','now'))",
		2, "Tolja", "Activnqj", "alfa@bravo.com", oneTwoThree, 0, 1, 1.15)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(INSERT) 2\n ", err)
		return
	}

	orderId := uuid.NewV4()
	_, err = DB.Exec("INSERT INTO orders (id, user_id, confirmed, date_created) VALUES (?,(SELECT id FROM users WHERE email = ?),?,strftime('%s','now'))",
		orderId, "alfa@bravo.com", 0)
	if err != nil {
		log.Println("sqlite.orders err \t", err)
	}

	_, err = DB.Exec("INSERT INTO users (id, firstname, lastname, email, password, is_admin, activated, user_percent, date_created) VALUES (?,?,?,?,?,?,?,?, strftime('%s','now'))",
		3, "Artem", "Non-active", "tema@bravo.com", oneTwoThree, 0, 0, 1.25)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(INSERT) 3\n ", err)
		return
	}

	_, err = DB.Exec("INSERT INTO users (id, firstname, lastname, email, password, is_admin, activated, user_percent, date_created) VALUES (?,?,?,?,?,?,?,?, strftime('%s','now'))",
		4, "Fedja", "PassiF", "fedos@bravo.com", oneTwoThree, 0, 0, 1.5)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(INSERT) 4\n ", err)
		return
	}
	log.Println("db created")
}
