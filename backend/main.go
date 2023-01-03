package main

import (
	"log"
	"net/http"
	"os"
	"pood/db"
)

func main() {
	// writing to file make it difficult to trouble shoot why conteiner wont go up
	// LOG_FILE := "./app_logs"
	// logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer logFile.Close()
	// log.SetOutput(logFile)
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// errorLog.SetOutput(logFile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// infoLog.SetOutput(logFile)

	setENV() // create .env file inside a container with all needed variables

	s := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  routes(),
	}

	infoLog.Println("Starting at http://localhost" + s.Addr)

	db.InitDatabase()
	defer db.DB.Close()

	err2 := s.ListenAndServe()
	if err2 != nil {
		errorLog.Fatal(err2)
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
