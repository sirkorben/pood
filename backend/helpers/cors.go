package helpers

import "net/http"

const (
	backend_ip               = "http://146.190.118.167" // for prod
	backend_ip_localhost     = "http://localhost"       // for running locally inside docker
	backend_ip_localhost3000 = "http://localhost:3000"  // for running locally
)

func HandleCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", backend_ip_localhost3000)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Authorization, Accept")
		h(w, r)
	}
}
