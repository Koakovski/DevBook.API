package controller

import (
	presenter "devbook-api/src/app/presenters"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UserFindOneController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.GetDbConnection()
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.GetUserRepository(db)

	user, err := userRepository.FindById(userId)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, user)
}
