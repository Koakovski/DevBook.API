package routes

import "net/http"

var UserRoutes = []Route{
	// CREATE USER ROUTE
	{
		Uri:       "/user",
		Method:    http.MethodPost,
		Handler:   func(http.ResponseWriter, *http.Request) {},
		IsPrivate: false,
	},
	// GET ALL USERS ROUTE
	{
		Uri:       "/user",
		Method:    http.MethodGet,
		Handler:   func(http.ResponseWriter, *http.Request) {},
		IsPrivate: false,
	},
	// GET ONE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodGet,
		Handler:   func(http.ResponseWriter, *http.Request) {},
		IsPrivate: false,
	},
	// UPDATE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodPut,
		Handler:   func(http.ResponseWriter, *http.Request) {},
		IsPrivate: false,
	},
	// DELETE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodDelete,
		Handler:   func(http.ResponseWriter, *http.Request) {},
		IsPrivate: false,
	},
}
