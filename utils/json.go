package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONWriter Write The JSON Response to The Response Writer
func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
