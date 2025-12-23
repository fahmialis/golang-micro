package service

import (
	"context"
	"fmt"
	"user-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]repository.User, error) {

	data, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("no users found")
	}

	return data, nil
}
