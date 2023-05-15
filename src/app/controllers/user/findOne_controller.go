package controller

import (
	presenter "devbook-api/src/app/presenters"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"net/http"
	"strconv"
)

func UserFindOneController(w http.ResponseWriter, r *http.Request) {
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

	user, err := userRepository.FindById(userId)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	if strconv.Itoa(int(user.ID)) != "" {
		presenter.ReponsePresenter(w, http.StatusBadRequest, nil)
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, user)
}
