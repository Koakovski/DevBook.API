package route

import (
	middlwares "devbook-api/src/app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent routes of API
type Route struct {
	Uri       string
	Method    string
	Handler   func(http.ResponseWriter, *http.Request)
	IsPrivate bool
}

func ConfigureRoutes(router *mux.Router) *mux.Router {
	routes := append(UserRoutes, AuthRoutes...)

	for _, route := range routes {
		if route.IsPrivate {
			router.HandleFunc(route.Uri, middlwares.Authenticate(route.Handler)).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
		}
	}

	return router
}
