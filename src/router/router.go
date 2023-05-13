package router

import "github.com/gorilla/mux"

// GenerateRouter returns a router with routes configured
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
