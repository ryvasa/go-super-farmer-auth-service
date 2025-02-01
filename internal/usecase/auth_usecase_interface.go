package usecase

import "context"

type AuthUsecase interface {
	GoogleLogin(ctx context.Context) (bool, error)
}
