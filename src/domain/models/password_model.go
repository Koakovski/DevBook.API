package model

import (
	"errors"
)

type Password struct {
	NewPassword     string `json:"newPassword"`
	CurrentPassword string `json:"currentPassword"`
}

func (password *Password) Validate() error {
	if password.CurrentPassword == "" {
		return errors.New("CurrentPassword is missing")
	}

	if password.NewPassword == "" {
		return errors.New("NewPassword is missing")
	}

	return nil
}
