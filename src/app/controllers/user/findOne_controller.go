package controller

import "net/http"

func UserFindOneController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encontrando um Usuário!"))
}
