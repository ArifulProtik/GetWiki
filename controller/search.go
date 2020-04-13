package controller

import (
	"getwiki/db"
	"getwiki/utils"
	"net/http"
	"net/url"
)

// SearchandGet Handles The Search and Get Wiki
func SearchandGet(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	prep := url.QueryEscape(param.Get("q"))
	articles, err := db.STRG.SearchArticle(prep)
	if err != nil {
		utils.JSONWriter(w, utils.H{
			"SearchText": param.Get("q"),
			"Error:":     err.Error(),
		}, http.StatusUnprocessableEntity)
	}
	if len(articles) != 0 {
		utils.JSONWriter(w, articles, 200)
	} else {
		get1, err := Getwiki(prep)
		if err != nil {
			utils.JSONWriter(w, utils.H{
				"SearchText": param.Get("q"),
				"Error:":     err.Error(),
			}, http.StatusUnprocessableEntity)
		} else {
			errs := db.STRG.CreateArticle(get1)
			if errs != nil {
				utils.JSONWriter(w, utils.H{
					"SearchText": param.Get("q"),
					"Error:":     err.Error(),
				}, http.StatusUnprocessableEntity)
			} else {
				getter, err := db.STRG.SearchArticle(prep)
				if err != nil {
					utils.JSONWriter(w, utils.H{
						"SearchText": param.Get("q"),
						"Error:":     err.Error(),
					}, http.StatusUnprocessableEntity)
				}
				utils.JSONWriter(w, getter, 200)
			}

		}
	}
}
