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

	// manage variables on the droplet in DO
	setENV()

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
	sm.HandleFunc("/api/me", helpers.HandleCors(me))                                         // GET
	sm.HandleFunc("/api/signup", helpers.HandleCors(signUp))                                 // POST
	sm.HandleFunc("/api/signin", helpers.HandleCors(signIn))                                 // POST
	sm.HandleFunc("/api/signout", helpers.HandleCors(signOut))                               // GET
	sm.HandleFunc("/api/search", helpers.HandleCors(search))                                 // POST
	sm.HandleFunc("/api/admin/", helpers.HandleCors(admin))                                  // GET
	sm.HandleFunc("/api/admin/approve", helpers.HandleCors(adminApproveHandler))             // GET, PATCH
	sm.HandleFunc("/api/admin/managepercent", helpers.HandleCors(adminIncreasePriceHandler)) // GET, PATCH
	sm.HandleFunc("/api/admin/orders", helpers.HandleCors(adminOrdersHandler))               // GET, PATCH (not implemented)  ???
	sm.HandleFunc("/api/admin/orders/order", helpers.HandleCors(adminOrderHandler))          // GET, PATCH (not implemented)  ???
	sm.HandleFunc("/api/myorders", helpers.HandleCors(userOrders))                           // GET
	sm.HandleFunc("/api/myorders/order", helpers.HandleCors(order))                          // GET by query parameter
	sm.HandleFunc("/api/cart", helpers.HandleCors(shoppingCart))                             // GET, POST
	sm.HandleFunc("/api/cart/add", helpers.HandleCors(addItemToCart))                        // POST
	sm.HandleFunc("/api/cart/confirm", helpers.HandleCors(confirmCart))                      // POST
	sm.HandleFunc("/pi/cart/remove", helpers.HandleCors(removeCart))                         // DELETE
	sm.HandleFunc("/api/cart/removeitem", helpers.HandleCors(removeItemFromCart))            // DELETE
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
