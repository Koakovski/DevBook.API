package controller

import (
	model "devbook-api/src/domain/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserFromBody(r *http.Request, isUpdate bool) (user model.User, err error, statusCode int) {
	var userModel model.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return userModel, err, http.StatusUnprocessableEntity
	}

	if err = json.Unmarshal(body, &userModel); err != nil {
		return userModel, err, http.StatusBadRequest
	}

	if err = userModel.Prepare(isUpdate); err != nil {
		return userModel, err, http.StatusBadRequest
	}

	return userModel, nil, 0
}

func GetUserId(r *http.Request) (uint64, error) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
