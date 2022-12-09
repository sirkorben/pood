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
	sm.HandleFunc("/", home)                              // GET
	sm.HandleFunc("/signup", signUp)                      // GET, POST
	sm.HandleFunc("/signin", signIn)                      // GET, POST
	sm.HandleFunc("/signout", signOut)                    // POST
	sm.HandleFunc("/search", search)                      // GET, POST
	sm.HandleFunc("/admin/", admin)                       // GET
	sm.HandleFunc("/admin/approve", adminApproveHandler)  // GET, PATCH
	sm.HandleFunc("/admin/orders", adminOrdersHandler)    // GET, PATCH(not implemented)
	sm.HandleFunc("/myorders", userOrders)                // GET
	sm.HandleFunc("/order", order)                        // GET
	sm.HandleFunc("/cart", shoppingCart)                  // GET, POST
	sm.HandleFunc("/cart/add", addItemToCart)             // POST
	sm.HandleFunc("/cart/confirm", confirmCart)           // POST
	sm.HandleFunc("/cart/remove", removeCart)             // DELETE
	sm.HandleFunc("/cart/removeitem", removeItemFromCart) // DELETE

	return sm
}
