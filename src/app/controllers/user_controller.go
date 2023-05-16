package controller

import (
	presenter "devbook-api/src/app/presenters"
	model "devbook-api/src/domain/models"
	"devbook-api/src/infra/auth"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories"
	"devbook-api/src/infra/security"
	"errors"
	"net/http"
	"strings"
)

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	var userModel model.User
	statusCode, err := GetDataFromBody(r, &userModel, false)
	if err != nil {
		presenter.ErrorPresenter(w, statusCode, err)
		return
	}

	if err = userModel.Prepare(false); err != nil {
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

func UserDeleteController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	_, statusCode, err := IsUserAllowed(r, userId)
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
	if err = userRepository.Delete(userId); err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusNoContent, nil)
}

func UserFindAllController(w http.ResponseWriter, r *http.Request) {
	nameOrNickName := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.GetDbConnection()
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.GetUserRepository(db)

	users, err := userRepository.FindAll(nameOrNickName)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, users)
}

func UserFindOneController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetId(r)
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

	if user.ID == 0 {
		presenter.ErrorPresenter(w, http.StatusBadRequest, errors.New("no user found"))
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, user)
}

func UserUpdateController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	var userModel model.User
	statusCode, err := GetDataFromBody(r, &userModel, true)
	if err != nil {
		presenter.ErrorPresenter(w, statusCode, err)
		return
	}

	_, statusCode, err = IsUserAllowed(r, userId)
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

func UserFollowUserController(w http.ResponseWriter, r *http.Request) {
	userToFollowId, err := GetId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	requestingUserId, err := auth.ExtractUserId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusUnauthorized, err)
		return
	}

	if userToFollowId == requestingUserId {
		presenter.ErrorPresenter(w, http.StatusForbidden, errors.New("user cannot follow himself"))
		return
	}

	db, err := database.GetDbConnection()
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.GetUserRepository(db)

	if err = userRepository.Follow(requestingUserId, userToFollowId); err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusNoContent, nil)
}

func UserUnfollowUserController(w http.ResponseWriter, r *http.Request) {
	userToUnfollowId, err := GetId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	requestingUserId, err := auth.ExtractUserId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusUnauthorized, err)
		return
	}

	if userToUnfollowId == requestingUserId {
		presenter.ErrorPresenter(w, http.StatusForbidden, errors.New("user cannot unfollow himself"))
		return
	}

	db, err := database.GetDbConnection()
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.GetUserRepository(db)

	if err = userRepository.Unfollow(requestingUserId, userToUnfollowId); err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusNoContent, nil)
}

func UserFindAllUserFollowersController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetId(r)
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

	followers, err := userRepository.FindAllUserFollowers(userId)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, followers)
}

func UserFindAllUserFollowingController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetId(r)
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

	followers, err := userRepository.FindAllUserFollowing(userId)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, followers)
}

func UserUpdatePasswordController(w http.ResponseWriter, r *http.Request) {
	userId, err := GetId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	requestingUserId, statusCode, err := IsUserAllowed(r, userId)
	if err != nil {
		presenter.ErrorPresenter(w, statusCode, err)
		return
	}

	var passwordModel model.Password
	statusCode, err = GetDataFromBody(r, &passwordModel, true)
	if err != nil {
		presenter.ErrorPresenter(w, statusCode, err)
		return
	}

	if err := passwordModel.Validate(); err != nil {
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
	databaseUser, err := userRepository.FindById(requestingUserId)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.ComparePassword(databaseUser.Password, passwordModel.CurrentPassword); err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, errors.New("current password don't match"))
		return
	}

	newPasswordHashed, err := security.HashPassword(passwordModel.NewPassword)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	if err := userRepository.UpdatePassword(string(newPasswordHashed), requestingUserId); err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusNoContent, nil)
}
