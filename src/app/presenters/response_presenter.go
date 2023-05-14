package presenter

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReponsePresenter(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
