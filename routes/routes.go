// Routes package for defining all the routes in our app
package routes

import (
	"getwiki/controller"
	"getwiki/handler"

	"github.com/gorilla/mux"
)

// RegisterRouter Regiosters All The Router
func RegisterRouter(r *mux.Router) {
	apiv1 := r.PathPrefix("/api/v1").Subrouter()
	// Publicly Accesseble Endpoint goes here
	apiv1.Handle("/", controller.BasicMiddleWare(handler.NotImplementedHandler))
}
