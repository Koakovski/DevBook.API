package controller

import (
	presenter "devbook-api/src/app/presenters"
	model "devbook-api/src/domain/models"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"net/http"
)

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	var userModel model.User
	statusCode, err := GetDataFromBody(r, userModel, false)
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
	userModel.ID, err = userRepository.Create(userModel)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	userModel.Password = ""

	presenter.ReponsePresenter(w, http.StatusCreated, userModel)
}
