package route

import (
	controller "devbook-api/src/app/controllers/user"
	"net/http"
)

var UserRoutes = []Route{
	// CREATE USER ROUTE
	{
		Uri:       "/user",
		Method:    http.MethodPost,
		Handler:   controller.UserCreateController,
		IsPrivate: false,
	},
	// FIND ALL USERS ROUTE
	{
		Uri:       "/user",
		Method:    http.MethodGet,
		Handler:   controller.UserFindAllController,
		IsPrivate: false,
	},
	// FIND ONE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodGet,
		Handler:   controller.UserFindOneController,
		IsPrivate: false,
	},
	// UPDATE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodPut,
		Handler:   controller.UserUpdateController,
		IsPrivate: false,
	},
	// DELETE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodDelete,
		Handler:   controller.UserDeleteController,
		IsPrivate: false,
	},
}
