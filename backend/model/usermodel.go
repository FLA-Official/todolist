package model

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

type User struct {
	ID        int       `db:"id" json:"userid"`
	Username  string    `db:"user_name" json:"username"`
	Fullname  string    `db:"full_name" json:"fullname"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (u *User) Validate() error {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

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
