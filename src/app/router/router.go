package router

import (
	"devbook-api/src/app/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter returns a router with routes configured
func GenerateRouter() *mux.Router {
	return routes.ConfigureRoutes(mux.NewRouter())
}
