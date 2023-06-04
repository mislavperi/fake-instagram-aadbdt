package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Whoami() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *UserController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
