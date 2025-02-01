package repository

import (
	"context"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db}
}

func (a *AuthRepositoryImpl) GoogleLogin(ctx context.Context, user string) (bool, error) {
	return true, nil
}
