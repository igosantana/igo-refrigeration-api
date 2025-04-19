package user

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers(ctx context.Context) ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
