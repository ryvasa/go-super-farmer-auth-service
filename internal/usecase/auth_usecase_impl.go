package usecase

import "context"

type AuthUsecaseImpl struct {
}

func NewAuthUsecaseImpl() AuthUsecase {
	return &AuthUsecaseImpl{}
}

func (a *AuthUsecaseImpl) GoogleLogin(ctx context.Context) (bool, error) {
	return true, nil
}
