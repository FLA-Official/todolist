package model

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

type User struct {
	ID        int       `json:"userid"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

func (u *User) Validate() error {
	u.Username = strings.TrimSpace(u.Username)
	u.Fullname = strings.TrimSpace(u.Fullname)
	u.Email = strings.TrimSpace(u.Email)
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

	if u.Email == "" {
		return errors.New("email is required")
	}

	// simple email validation
	err := validateEmail(u.Email)
	if err != nil {
		return err
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email format")
	}
	return nil
}
