package middleware

import (
	"log"
	"net/http"
)

//TokenValidationMiddleware app
func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Accessing : %v", r.RequestURI)
		// http.Error(w,"Middleware Error", http.StatusInternalServerError)
	})
}
