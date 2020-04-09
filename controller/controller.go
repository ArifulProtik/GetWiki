// package for controlling Requests
package controller

import (
	"net/http"
)

// BasicMiddleWare Server The Basic MiddleWare Utilities
func BasicMiddleWare(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json; charset=UTF8")
		next(w, r)
	})
}
