//go:build wireinject
// +build wireinject

package wire_auth

import (
	"github.com/google/wire"
	"github.com/ryvasa/go-super-farmer-auth-service/cmd/app"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/delivery/http/handler"
	route "github.com/ryvasa/go-super-farmer-auth-service/internal/delivery/http/routes"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/repository"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/database"
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/oauth"
)

var allSet = wire.NewSet(
	// Infrastructure
	env.LoadEnv,
	oauth.NewOauth,
	database.NewPostgres,

	// Repository
	repository.NewAuthRepositoryImpl,

	// Service
	usecase.NewAuthUsecaseImpl,

	// Handler
	handler.NewAuthHandlerImpl,

	// App
	app.NewApp,

	route.NewRoutes,
	handler.NewHandlers,
)

func InitializeAuthApp() (*app.AuthApp, error) {
	wire.Build(allSet)
	return nil, nil
}
