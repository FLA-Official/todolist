package model

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        int       `json:"userid"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Gmail     string    `json:"gmail"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

func (u *User) Validate() error {
	u.Username = strings.TrimSpace(u.Username)
	u.Fullname = strings.TrimSpace(u.Fullname)
	u.Gmail = strings.TrimSpace(u.Gmail)
	u.Password = strings.TrimSpace(u.Password)

	if u.Username == "" {
		return errors.New("username is required")
	}

	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	if strings.Contains(u.Username, " ") {
		return errors.New("username cannot contain spaces")
	}

	if u.Fullname == "" {
		return errors.New("fullname is required")
	}

	if u.Gmail == "" {
		return errors.New("email is required")
	}

	// simple email validation
	if !strings.Contains(u.Gmail, "@") || !strings.Contains(u.Gmail, ".") {
		return errors.New("invalid email format")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}
