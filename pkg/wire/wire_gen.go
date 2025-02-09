// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire_auth

import (
	"github.com/google/wire"
	"github.com/ryvasa/go-super-farmer-auth-service/cmd/app"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/delivery/http/handler"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/delivery/http/routes"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/repository"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/database"
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/oauth"
)

// Injectors from wire.go:

func InitializeAuthApp() (*app.AuthApp, error) {
	envEnv, err := env.LoadEnv()
	if err != nil {
		return nil, err
	}
	config := oauth.NewOauth(envEnv)
	authHandler := handler.NewAuthHandlerImpl(config)
	handlers := handler.NewHandlers(authHandler)
	engine := route.NewRoutes(handlers)
	db, err := database.NewPostgres(envEnv)
	if err != nil {
		return nil, err
	}
	authApp := app.NewApp(engine, envEnv, db)
	return authApp, nil
}

// wire.go:

var allSet = wire.NewSet(env.LoadEnv, oauth.NewOauth, database.NewPostgres, repository.NewAuthRepositoryImpl, usecase.NewAuthUsecaseImpl, handler.NewAuthHandlerImpl, app.NewApp, route.NewRoutes, handler.NewHandlers)
