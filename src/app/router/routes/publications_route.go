package route

import (
	controller "devbook-api/src/app/controllers"
	"net/http"
)

var PublicationRoute = []Route{
	// CREATE A PUBLICATION
	{
		Uri:       "/publication",
		Method:    http.MethodPost,
		Handler:   controller.PublicationCreateController,
		IsPrivate: false,
	},
	// FIND ALL PUBLICATIONS
	{
		Uri:    "/publication",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("FIND ALL PUBLICATIONS"))
		},
		IsPrivate: false,
	},
	// FIND ONE PUBLICATION
	{
		Uri:       "/publication/{id}",
		Method:    http.MethodGet,
		Handler:   controller.PublicationFindOneController,
		IsPrivate: false,
	},
	// UPDATE PUBLICATION
	{
		Uri:    "/publication/{id}",
		Method: http.MethodPut,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("UPDATE PUBLICATION"))
		},
		IsPrivate: false,
	},
	// DELETE PUBLICATION
	{
		Uri:    "/publication/{id}",
		Method: http.MethodDelete,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("DELETE PUBLICATION"))
		},
		IsPrivate: false,
	},
}