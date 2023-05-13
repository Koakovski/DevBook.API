package routes

import (
	controllers "devbook-api/src/app/controllers/user"
	"net/http"
)

var UserRoutes = []Route{
	// CREATE USER ROUTE
	{
		Uri:       "/user",
		Method:    http.MethodPost,
		Handler:   controllers.UserCreateController,
		IsPrivate: false,
	},
	// FIND ALL USERS ROUTE
	{
		Uri:       "/user",
		Method:    http.MethodGet,
		Handler:   controllers.UserFindAllController,
		IsPrivate: false,
	},
	// FIND ONE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodGet,
		Handler:   controllers.UserFindOneController,
		IsPrivate: false,
	},
	// UPDATE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodPut,
		Handler:   controllers.UserUpdateController,
		IsPrivate: false,
	},
	// DELETE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodDelete,
		Handler:   controllers.UserDeleteController,
		IsPrivate: false,
	},
}
