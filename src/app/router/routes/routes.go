package route

import (
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
	routes := UserRoutes

	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	return router
}
