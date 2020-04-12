package controller

import (
	"getwiki/utils"
	"net/http"
)

// BasicMiddleWare Server The Basic MiddleWare Utilities
func BasicMiddleWare(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json; charset=UTF8")
		next(w, r)
	})
}

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	utils.JSONWriter(w, utils.H{
		"msg":    "It is Not Implemented",
		"Status": http.StatusOK,
	}, 200)
}
