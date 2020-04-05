// handler for Returing Response to the Handler
package handler

import (
	"getwiki/utils"
	"net/http"
)

// NotImplementedHandler retuns the Not Implemented Messege.
func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	utils.JSONWriter(w, utils.H{
		"msg":    "It is Not Implemented",
		"Status": http.StatusOK,
	}, 200)
}
