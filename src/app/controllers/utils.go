package controller

import (
	"encoding/json"
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

func GetUserId(r *http.Request) (uint64, error) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
