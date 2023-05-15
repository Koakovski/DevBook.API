package controller

import (
	presenter "devbook-api/src/app/presenters"
	model "devbook-api/src/domain/models"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"devbook-api/src/infra/security"
	"net/http"
)

func AuthLoginController(w http.ResponseWriter, r *http.Request) {
	var userModel model.User
	statusCode, err := GetDataFromBody(r, userModel, false)
	if err != nil {
		presenter.ErrorPresenter(w, statusCode, err)
		return
	}

	db, err := database.GetDbConnection()
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.GetUserRepository(db)

	databaseUser, err := userRepository.FindByEmail(userModel.Email)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ComparePassword(userModel.Password, databaseUser.Password); err != nil {
		presenter.ErrorPresenter(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte("Logado"))
}
