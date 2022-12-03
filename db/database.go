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

	adminPass, err := helpers.GeneratePassword("dummyPassWillBeTakenFromLocalEnvOrSomeHow")
	if err != nil {
		// handle error
		return
	}
	const CREATE string = `
	create table users(	
		id INTEGER not null primary key autoincrement,
		firstname TEXT not null,
		lastname TEXT not null,
		email TEXT not null unique, 
		password BLOB not null,
		is_admin INTEGER not null,
		activated INTEGER not null,
		date_created INTEGER not null);

	create table sessions(
		id TEXT not null primary key, 
		user_id INTEGER not null unique, 
		created_date INTEGER not null);
	`

	_, err = DB.Exec(CREATE)
	if err != nil {
		// handle error
		// error comes here while starting application with created database is in root
		log.Println("error in _, err = DB.Exec(CREATE)", err)
		return
	}

	_, err = DB.Exec("INSERT INTO users (id, firstname, lastname, email, password, is_admin, activated, date_created) VALUES (?,?,?,?,?,?,?, strftime('%s','now'))",
		1, "Daniil", "Batjkovich", "danic@prostoSobaka.com", adminPass, 1, 1)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(INSERT)\n ", err)
		return
	}

	// delete two users below
	// delete passwords
	oneTwoThree, err := helpers.GeneratePassword("123456")
	if err != nil {
		// handle error
		return
	}
	_, err = DB.Exec("INSERT INTO users (id, firstname, lastname, email, password, is_admin, activated, date_created) VALUES (?,?,?,?,?,?,?, strftime('%s','now'))",
		2, "User", "Activated", "alfa@bravo.com", oneTwoThree, 0, 1)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(INSERT)\n ", err)
		return
	}

	_, err = DB.Exec("INSERT INTO users (id, firstname, lastname, email, password, is_admin, activated, date_created) VALUES (?,?,?,?,?,?,?, strftime('%s','now'))",
		3, "Artem", "Bulavin", "bula@bravo.com", oneTwoThree, 0, 0)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(INSERT)\n ", err)
		return
	}
}
