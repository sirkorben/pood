//TODO 1: make sure that database are created properly with two tables for now
//TODO 2: insert dummy user to DB
//TODO 3: setup proper logging
//TODO 3: handle errors properly with logging regarding db errors
//TODO later: think about possibility to use real db if it is needed ?

//HINTS: probably while hosting app we need to set up environment and go get github.com/mattn/go-sqlite3

package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB = &sql.DB{}

func InitDatabase() {
	createConnection()
	createTables()
}

func createConnection() {
	db, err := sql.Open("sqlite3", "pood.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = db
}

func createTables() {
	const CREATE string = `
	create table users(	
		id INTEGER not null primary key autoincrement,
		firstname TEXT not null,
		lastname TEXT not null,
		age INTEGER not null,
		gender TEXT not null,
		username TEXT not null unique,
		email TEXT not null unique, 
		password BLOB not null, 
		creation_date INTEGER not null );

	create table sessions(
		id TEXT not null primary key, 
		user_id INTEGER not null unique, 
		created_date INTEGER not null);

	`

	// const INSERT string = `
	// INSERT INTO categories (id, name) VALUES (1, 'Cars');

	// `

	var err error

	_, err = DB.Exec(CREATE)
	if err != nil {
		// handle error
		log.Println("error in _, err = DB.Exec(CREATE)")

		return
	}

	// _, err = DB.Exec(INSERT)
	// if err != nil {
	// 	// handle error
	// 	log.Println("error in _, err = DB.Exec(INSERT)")
	// 	return
	// }

	if err = DB.Ping(); err != nil {
		// handle error
		log.Fatal(err)
	}
}
