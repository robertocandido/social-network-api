package models

import (
	"errors"
	"strings"
	"time"
)

// User it is a user for the social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Pass      string    `json:"pass,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare validate and format the user to persist
func (user *User) Prepare() error {
	if error := user.validate(); error != nil {
		return error
	}
	user.formater()

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The name is required field")
	}
	if user.Nick == "" {
		return errors.New("The nick is required field")
	}
	if user.Email == "" {
		return errors.New("The email is required field")
	}
	if user.Pass == "" {
		return errors.New("The pass is required field")
	}

	return nil
}

func (user *User) formater() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
