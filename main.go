package main

import (
	"log"
	"net/http"
	"os"
	"pood/db"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	s := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  routes(),
	}
	infoLog.Println("Starting at http://localhost" + s.Addr)

	db.InitDatabase()
	defer db.DB.Close()

	err := s.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func routes() *http.ServeMux {
	sm := http.NewServeMux()
	sm.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	sm.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	sm.HandleFunc("/", home)
	sm.HandleFunc("/signup", signUp)
	sm.HandleFunc("/signin", signIn)
	sm.HandleFunc("/signout", signOut)
	sm.HandleFunc("/search", search)
	return sm
}
