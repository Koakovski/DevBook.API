package controller

import (
	"devbook-api/src/infra/auth"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetDataFromBody(r *http.Request, value any, isUpdate bool) (statusCode int, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return http.StatusUnprocessableEntity, err
	}

	if err = json.Unmarshal(body, &value); err != nil {
		return http.StatusBadRequest, err
	}

	return 0, nil
}

func GetId(r *http.Request) (uint64, error) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func IsUserAllowed(r *http.Request, affectedUserId uint64) (userId uint64, statusCode int, err error) {
	requestingUserId, err := auth.ExtractUserId(r)
	if err != nil {
		return 0, http.StatusUnauthorized, err
	}
	if requestingUserId != affectedUserId {
		return 0, http.StatusForbidden, errors.New("not allowed")
	}

	return requestingUserId, 0, nil
}
