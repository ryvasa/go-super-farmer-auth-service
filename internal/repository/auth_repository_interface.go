package repository

import "context"

type AuthRepository interface {
	GoogleLogin(ctx context.Context, user string) (bool, error)
}
