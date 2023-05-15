package controller

import (
	presenter "devbook-api/src/app/presenters"
	model "devbook-api/src/domain/models"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"net/http"
)

func UserUpdateController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	var userModel model.User
	statusCode, err := GetDataFromBody(r, userModel, true)
	if err != nil {
		presenter.ErrorPresenter(w, statusCode, err)
		return
	}

	if err = userModel.Prepare(true); err != nil {
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

	if err = userRepository.Update(userId, userModel); err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusNoContent, nil)
}
