package controllers

import "net/http"

func UserFindAllController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encontrando todos os Usuário!"))
}
