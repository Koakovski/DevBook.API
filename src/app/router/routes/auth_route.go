package route

import (
	controller "devbook-api/src/app/controllers"
	"net/http"
)

var AuthRoutes = []Route{
	// LOGIN
	{
		Uri:       "/login",
		Method:    http.MethodPost,
		Handler:   controller.AuthLoginController,
		IsPrivate: false,
	},
}
