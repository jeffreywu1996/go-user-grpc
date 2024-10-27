package service

import (
	"context"

	"github.com/jeffreywu1996/go-user/internal/model"
	"github.com/jeffreywu1996/go-user/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	return s.repo.Create(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}
