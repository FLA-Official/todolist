package repo

import (
	"errors"
	"todolist/model"
	"todolist/utils"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserRepo interface {
	CreateUser(user *model.User) error
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
	ListUsers() ([]model.User, error)
	Find(email string) (*model.User, error)
}

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
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	query := `
	INSERT INTO users (user_name, full_name, email, password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at;
	`

	return u.dbCon.QueryRow(
		query,
		user.Username,
		user.Fullname,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (u *userRepo) GetUserByID(id int) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE id=$1`

	err := u.dbCon.Get(&user, query, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return nil, errors.New("username or email already exists")
			}
		}
		return nil, err
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

func (u *userRepo) Find(email string) (*model.User, error) {
	var user model.User

	query := `
	SELECT * FROM users 
	WHERE email=$1 
	`

	err := u.dbCon.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("No user with this email")
	}

	return &user, nil
}
