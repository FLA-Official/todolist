package service

import (
	"errors"
	"todolist/model"
	"todolist/repo"
	"todolist/utils"
)

type UserService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(user *model.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) Login(email, password string) (*model.User, error) {

	user, err := s.userRepo.Find(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// compare password
	err = utils.CheckPassword(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(uUser *model.User) error {
	return s.userRepo.UpdateUser(uUser)
}
