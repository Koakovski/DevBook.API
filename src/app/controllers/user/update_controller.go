package controllers

import "net/http"

func UserUpdateController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um Usuário!"))
}
