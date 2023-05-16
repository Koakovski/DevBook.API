package route

import (
	controller "devbook-api/src/app/controllers"
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
		IsPrivate: true,
	},
	// FIND ONE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodGet,
		Handler:   controller.UserFindOneController,
		IsPrivate: true,
	},
	// UPDATE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodPut,
		Handler:   controller.UserUpdateController,
		IsPrivate: true,
	},
	// DELETE USER ROUTE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodDelete,
		Handler:   controller.UserDeleteController,
		IsPrivate: true,
	},
	// FOLLOW A USER
	{
		Uri:       "/user/{id}/follow",
		Method:    http.MethodPost,
		Handler:   controller.UserFollowUserController,
		IsPrivate: true,
	},
	// UNFOLLOW A USER
	{
		Uri:       "/user/{id}/unfollow",
		Method:    http.MethodPost,
		Handler:   controller.UserUnfollowUserController,
		IsPrivate: true,
	},
}
