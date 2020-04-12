package controller

import (
	"getwiki/db"
	"getwiki/utils"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// SearchandGet Handles The Search and Get Wiki
func SearchandGet(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	prep := url.QueryEscape(param["text"])
	articles, err := db.STRG.SearchArticle(prep)
	if err != nil {
		utils.JSONWriter(w, utils.H{
			"SearchText": param["text"],
			"Error:":     err.Error(),
		}, http.StatusUnprocessableEntity)
	}
	if len(articles) != 0 {
		utils.JSONWriter(w, articles, 200)
	} else {
		getter, err := Getwiki(prep)
		if err != nil {
			utils.JSONWriter(w, utils.H{
				"SearchText": param["text"],
				"Error:":     err.Error(),
			}, http.StatusUnprocessableEntity)
		} else {
			errs := db.STRG.CreateArticle(getter)
			if errs != nil {
				utils.JSONWriter(w, utils.H{
					"SearchText": param["text"],
					"Error:":     err.Error(),
				}, http.StatusUnprocessableEntity)
			} else {
				utils.JSONWriter(w, getter, 200)
			}

		}
	}
}
