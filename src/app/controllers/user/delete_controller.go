package controller

import (
	presenter "devbook-api/src/app/presenters"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"net/http"
)

func UserDeleteController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
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
	if err = userRepository.Delete(userId); err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusNoContent, nil)
}
