package controller

import (
	presenter "devbook-api/src/app/presenters"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"net/http"
)

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	user, err, statusCode := GetUserFromBody(r, false)
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
	user.ID, err = userRepository.Create(user)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = ""

	presenter.ReponsePresenter(w, http.StatusCreated, user)
}
