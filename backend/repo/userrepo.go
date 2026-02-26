package repo

import (
	"errors"
	"todolist/model"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(user *model.User) error
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
	ListUsers() ([]model.User, error)
	Find(email, password string) (*model.User, error)
}

// ID        int       `json:"userid"`
// Username  string    `json:"username"`
// Fullname  string    `json:"fullname"`
// Gmail     string    `json:"gmail"`
// Password  string    `json:"password"`
// Role      string    `json:"role"` // developer, superviser, Admin, viewer
// Createdat time.Time `json:"createdat"`

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (u *userRepo) CreateUser(user *model.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	query := `
	INSERT INTO users (username, fullname, gmail, password, role)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, createdat;
	`

	return u.dbCon.QueryRow(
		query,
		user.Username,
		user.Fullname,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)
}

func (u *userRepo) GetUserByID(id int) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE id=$1`

	err := u.dbCon.Get(&user, query, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *userRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE gmail=$1`

	err := u.dbCon.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("no user found with this email")
	}

	return &user, nil
}

func (u *userRepo) UpdateUser(user *model.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	query := `
	UPDATE users
	SET username=$1, fullname=$2, gmail=$3, password=$4
	WHERE id=$5
	`

	result, err := u.dbCon.Exec(
		query,
		user.Username,
		user.Fullname,
		user.Email,
		user.Password,
		user.ID,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (u *userRepo) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`

	result, err := u.dbCon.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (u *userRepo) ListUsers() ([]model.User, error) {
	var users []model.User

	query := `SELECT * FROM users ORDER BY id`

	err := u.dbCon.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepo) Find(email, password string) (*model.User, error) {
	var user model.User

	query := `
	SELECT * FROM users 
	WHERE gmail=$1 AND password=$2
	`

	err := u.dbCon.Get(&user, query, email, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
