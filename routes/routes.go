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
	apiv1.Handle("/search", controller.BasicMiddleWare(controller.SearchandGet)).Queries("q", "{q}").Methods("GET")
	apiv1.Handle("/posts", controller.BasicMiddleWare(controller.Allposts)).Methods("GET")
	apiv1.Handle("/post/{id}", controller.BasicMiddleWare(controller.Getpost)).Methods("GET")
}
