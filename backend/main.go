package main

import (
	"log"
	"net/http"
	"os"
	"pood/db"
)

func main() {

	//logging ??
	// log to custom file
	LOG_FILE := "./testlogfile"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Logging to custom file")
	//logging
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.SetOutput(logFile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.SetOutput(logFile)

	s := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  routes(),
	}
	infoLog.Println("Starting at http://localhost" + s.Addr)

	db.InitDatabase()
	defer db.DB.Close()

	err2 := s.ListenAndServe()
	if err != nil {
		errorLog.Println("I am from Alex")
		log.Fatal(err2)
	}
}

func routes() *http.ServeMux {
	sm := http.NewServeMux()
	sm.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	sm.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	sm.HandleFunc("/", home)                              // GET
	sm.HandleFunc("/signup", signUp)                      // GET, POST
	sm.HandleFunc("/signin", signIn)                      // GET, POST
	sm.HandleFunc("/signout", signOut)                    // POST
	sm.HandleFunc("/search", search)                      // GET, POST
	sm.HandleFunc("/admin/", admin)                       // GET
	sm.HandleFunc("/admin/approve", adminApproveHandler)  // GET, PATCH
	sm.HandleFunc("/admin/orders", adminOrdersHandler)    // GET, PATCH(not implemented)
	sm.HandleFunc("/myorders", userOrders)                // GET
	sm.HandleFunc("/order", order)                        // GET by query parameter
	sm.HandleFunc("/cart", shoppingCart)                  // GET, POST
	sm.HandleFunc("/cart/add", addItemToCart)             // POST
	sm.HandleFunc("/cart/confirm", confirmCart)           // POST
	sm.HandleFunc("/cart/remove", removeCart)             // DELETE
	sm.HandleFunc("/cart/removeitem", removeItemFromCart) // DELETE

	return sm
}
