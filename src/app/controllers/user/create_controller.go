package controller

import (
	presenter "devbook-api/src/app/presenters"
	model "devbook-api/src/domain/models"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"encoding/json"
	"io"
	"net/http"
)

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
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
	user.ID, err = userRepository.Create(user)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = ""

	presenter.ReponsePresenter(w, http.StatusCreated, user)
}
