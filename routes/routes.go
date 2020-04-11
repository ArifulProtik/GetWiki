// Routes package for defining all the routes in our app
package routes

import (
	"getwiki/controller"

	"github.com/gorilla/mux"
)

// RegisterRouter Regiosters All The Router
func RegisterRouter(r *mux.Router) {
	apiv1 := r.PathPrefix("/api/v1").Subrouter()
	// Publicly Accesseble Endpoint goes here
	apiv1.Handle("/", controller.BasicMiddleWare(controller.NotImplementedHandler))
	apiv1.Handle("/search/{text}", controller.BasicMiddleWare(controller.SearchandGet)).Methods("GET")
}
