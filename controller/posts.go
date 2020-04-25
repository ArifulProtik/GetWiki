package controller

import (
	"getwiki/db"
	"getwiki/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Allposts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	posts, err := db.STRG.Getall()
	if err != nil {
		utils.JSONWriter(w, utils.H{
			"Error": err.Error(),
		}, http.StatusUnprocessableEntity)
	} else {
		utils.JSONWriter(w, posts, 200)
	}
}
func Getpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(r)
	id, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		utils.JSONWriter(w, utils.H{
			"Error": err.Error(),
		}, http.StatusBadRequest)
	}
	post, err := db.STRG.Getbyid(id)
	if err != nil {
		utils.JSONWriter(w, utils.H{
			"Error": err.Error(),
		}, http.StatusUnprocessableEntity)
	} else {
		utils.JSONWriter(w, post, 200)
	}
}
