package routes

import "net/http"

// Route represent routes of API
type Route struct {
	Uri       string
	Method    string
	Handler   func(http.ResponseWriter, *http.Request)
	IsPrivate bool
}
