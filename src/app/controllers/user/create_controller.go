package controllers

import "net/http"

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}
