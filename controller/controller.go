// package for controlling Requests
package controller

import (
	"log"
	"net/http"
)

// BasicMiddleWare Server The Basic MiddleWare Utilities
func BasicMiddleWare(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello From Middleware")
		next(w, r)
	})
}
