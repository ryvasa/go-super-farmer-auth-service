package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/env"
	"gorm.io/gorm"
)

type AuthApp struct {
	Router *gin.Engine
	Env    *env.Env
	DB     *gorm.DB
	// Handler  *report_handler.Handlers
}

func NewApp(
	router *gin.Engine,
	env *env.Env,
	db *gorm.DB,
	// handler *report_handler.Handlers,
) *AuthApp {
	return &AuthApp{
		Router: router,
		Env:    env,
		DB:     db,
		// Handler:  handler,
	}
}
