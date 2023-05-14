package presenter

import "net/http"

func ErrorPresenter(w http.ResponseWriter, statusCode int, err error) {
	ReponsePresenter(w, statusCode, struct {
		Error string `json:"erro"`
	}{
		Error: err.Error(),
	})
}
