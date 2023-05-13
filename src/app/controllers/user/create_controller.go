package controller

import (
	model "devbook-api/src/domain/models"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories/user"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := repository.GetUserRepository(db)
	createdUserId, err := userRepository.Create(user)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%d", createdUserId)))
}
