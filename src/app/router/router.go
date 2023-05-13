package router

import (
	route "devbook-api/src/app/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter returns a router with routes configured
func GenerateRouter() *mux.Router {
	return route.ConfigureRoutes(mux.NewRouter())
}
