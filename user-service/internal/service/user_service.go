package service

import (
	"context"
	"fmt"
	db "user-service/internal/db/gen"

	"github.com/google/uuid"
)

type UserService struct {
	q *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{q: q}
}

func (s *UserService) GetUsers(ctx context.Context) ([]db.LocalProfile, error) {
	users, err := s.q.ListLocalProfiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*db.LocalProfile, error) {

	fmt.Println("id", id)
	user, err := s.q.GetLocalProfileByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by ID: %w", err)
	}
	return &user, nil
}
