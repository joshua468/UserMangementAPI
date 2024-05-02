package service

import (
	"context"
	"errors"

	"github.com/joshua468/usermanagementapi/model"
	"github.com/joshua468/usermanagementapi/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
