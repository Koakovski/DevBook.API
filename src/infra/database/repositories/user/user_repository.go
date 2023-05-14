package repository

import (
	"database/sql"
	model "devbook-api/src/domain/models"
)

type userRepository struct {
	db *sql.DB
}

func GetUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (userRepository userRepository) Create(user model.User) (uint64, error) {
	statment, err := userRepository.db.Prepare(
		"INSERT INTO users (name, nickName, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.NickName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userId), nil
}
