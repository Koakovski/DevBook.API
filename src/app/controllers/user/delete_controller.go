package controller

import "net/http"

func UserDeleteController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
