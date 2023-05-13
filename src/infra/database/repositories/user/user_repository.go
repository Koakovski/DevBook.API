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

func (u userRepository) Create(user model.User) (uint64, error) {
	return 0, nil
}
