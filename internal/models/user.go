package models

import (
	"errors"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) CheckCredentials() error {
	if u.Username == "" || u.Password == "" {
		return errors.New("username or password invalid")
	}

	return nil
}
