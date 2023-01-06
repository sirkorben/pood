package main

import (
	"log"
	"net/http"
	"os"
	"pood/db"
	"pood/helpers"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	setENV() // create .env file inside a container with all needed variables

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
	sm.HandleFunc("/me", helpers.HandleCors(me))                                         // GET
	sm.HandleFunc("/signup", helpers.HandleCors(signUp))                                 // POST
	sm.HandleFunc("/signin", helpers.HandleCors(signIn))                                 // POST
	sm.HandleFunc("/signout", helpers.HandleCors(signOut))                               // GET
	sm.HandleFunc("/search", helpers.HandleCors(search))                                 // POST
	sm.HandleFunc("/admin/", helpers.HandleCors(admin))                                  // GET
	sm.HandleFunc("/admin/approve", helpers.HandleCors(adminApproveHandler))             // GET, PATCH
	sm.HandleFunc("/admin/managepercent", helpers.HandleCors(adminIncreasePriceHandler)) // GET, PATCH
	sm.HandleFunc("/admin/orders", helpers.HandleCors(adminOrdersHandler))               // GET, PATCH (not implemented)
	sm.HandleFunc("/myorders", helpers.HandleCors(userOrders))                           // GET
	sm.HandleFunc("/order", helpers.HandleCors(order))                                   // GET by query parameter
	sm.HandleFunc("/cart", helpers.HandleCors(shoppingCart))                             // GET, POST
	sm.HandleFunc("/cart/add", helpers.HandleCors(addItemToCart))                        // POST
	sm.HandleFunc("/cart/confirm", helpers.HandleCors(confirmCart))                      // POST
	sm.HandleFunc("/cart/remove", helpers.HandleCors(removeCart))                        // DELETE
	sm.HandleFunc("/cart/removeitem", helpers.HandleCors(removeItemFromCart))            // DELETE
	return sm
}

// delete code if not needed

// writing to file make it difficult to trouble shoot why conteiner wont go up
// LOG_FILE := "./app_logs"
// logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
// if err != nil {
// 	log.Panic(err)
// }
// defer logFile.Close()
// log.SetOutput(logFile)
// log.SetFlags(log.Lshortfile | log.LstdFlags)
// errorLog.SetOutput(logFile)
// infoLog.SetOutput(logFile)

// sm.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
// sm.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
