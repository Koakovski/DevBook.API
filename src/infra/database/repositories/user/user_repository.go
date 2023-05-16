package repository

import (
	"database/sql"
	model "devbook-api/src/domain/models"
	"fmt"
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

func (userRepository userRepository) FindAll(nameOrNickName string) ([]model.User, error) {
	nameOrNickName = fmt.Sprintf("%%%s%%", nameOrNickName)

	var users []model.User
	rows, err := userRepository.db.Query(
		"SELECT id, name, nickName, email, createdAt FROM users WHERE name LIKE ? OR nickName LIKE ?",
		nameOrNickName,
		nameOrNickName,
	)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userRepository userRepository) FindById(userId uint64) (model.User, error) {
	var user model.User

	rows, err := userRepository.db.Query(
		"SELECT id, name, nickName, email, createdAt FROM users WHERE id = ?", userId,
	)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt); err != nil {
			return user, err
		}
	}

	return user, nil
}

func (userRepository userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	rows, err := userRepository.db.Query(
		"SELECT id, password FROM users WHERE email = ?", email,
	)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Password); err != nil {
			return user, err
		}
	}

	return user, nil
}

func (userRepository userRepository) Update(userId uint64, user model.User) error {
	statment, err := userRepository.db.Prepare(
		"UPDATE users SET name = ?, nickName = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err = statment.Exec(user.Name, user.NickName, user.Email, userId); err != nil {
		return err
	}

	return nil
}

func (userRepository userRepository) Delete(userId uint64) error {
	statement, err := userRepository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId); err != nil {
		return err
	}

	return nil
}

func (userRepository userRepository) Follow(userId, userToFollowId uint64) error {
	statement, err := userRepository.db.Prepare(
		"INSERT IGNORE INTO followers (userId, followerId) VALUES (?,?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, userToFollowId); err != nil {
		return err
	}

	return nil
}

func (userRepository userRepository) Unfollow(userId, userToFollowId uint64) error {
	statement, err := userRepository.db.Prepare(
		"DELETE FROM followers WHERE userId = ? AND followerID = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, userToFollowId); err != nil {
		return err
	}

	return nil
}

func (userRepository userRepository) FindAllUserFollowers(userId uint64) ([]model.User, error) {
	var users []model.User

	rows, err := userRepository.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email, u.createdAt	FROM users u 
		INNER JOIN followers f ON u.id = f.followerId
		WHERE f.userId = ?
	`, userId)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userRepository userRepository) FindAllUserFollowing(userId uint64) ([]model.User, error) {
	var users []model.User

	rows, err := userRepository.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email, u.createdAt	FROM users u 
		INNER JOIN followers f ON u.id = f.userId
		WHERE f.followerId = ?
	`, userId)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userRepository userRepository) UpdatePassword(password string, userId uint64) error {
	statment, err := userRepository.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err = statment.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
