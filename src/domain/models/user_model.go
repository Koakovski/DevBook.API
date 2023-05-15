package model

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(isUpdate bool) error {
	if err := user.validate(isUpdate); err != nil {
		return err
	}

	user.format()

	return nil
}

func (user *User) validate(isUpdate bool) error {
	if user.Name == "" {
		return errors.New("name is missing")
	}

	if user.Email == "" {
		return errors.New("email is missing")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email")
	}

	if user.NickName == "" {
		return errors.New("nickName is missing")
	}

	if user.Password == "" && !isUpdate {
		return errors.New("password is missing")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.NickName = strings.TrimSpace(user.NickName)
}
