package model

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("name is missing")
	}

	if user.Email == "" {
		return errors.New("email is missing")
	}

	if user.NickName == "" {
		return errors.New("nickName is missing")
	}

	if user.Password == "" {
		return errors.New("password is missing")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.NickName = strings.TrimSpace(user.NickName)
}
