package handler

import "github.com/gin-gonic/gin"

type AuthHandlerImpl struct {
}

func NewAuthHandlerImpl() AuthHandler {
	return &AuthHandlerImpl{}
}

func (a *AuthHandlerImpl) GoogleLogin(ctx *gin.Context) {
}
