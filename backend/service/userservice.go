package service

import (
	"context"
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

func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {

	logger := utils.LoggerFromContext(ctx)

	user, err := s.userRepo.Find(email)
	if err != nil {
		logger.Error("user not found", "email", email)
		return nil, errors.New("invalid credentials")
	}

	err = utils.CheckPassword(user.Password, password)
	if err != nil {
		logger.Error("invalid password", "email", email)
		return nil, errors.New("invalid credentials")
	}

	logger.Info("user logged in", "user_id", user.ID)

	return user, nil
}

func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(uUser *model.User) error {
	return s.userRepo.UpdateUser(uUser)
}
