package user

import "context"

type Service interface {
	GetAllUsers(ctx context.Context) ([]User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllUsers(ctx context.Context) ([]User, error) {
	return s.repo.GetAllUsers(ctx)
}
