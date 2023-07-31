package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	customerrors "github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
)

type UserService interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	Login(username string, password string) (*string, *string, error)
	GetUserInformation(token string) (*models.User, error)
}

type UserController struct {
	UserService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) Whoami() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "turn around")
	}
}

func (c *UserController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *UserController) RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser models.User
		if err := ctx.BindJSON(&newUser); err != nil {
			return
		}
		err := c.UserService.Create(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Email, newUser.Password)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, errors.New("error while registering user"))
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *UserController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, nil)
			return
		}
		accessToken, refreshToken, err := c.UserService.Login(user.Username, user.Password)
		if err != nil {
			if customerrors.IsInvalidCredentialsError(err) {
				fmt.Println(err)
				ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid credentials"))
				return
			}
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.SetCookie("accessToken", *accessToken, 3600, "", "localhost", false, false)
		ctx.SetCookie("refreshToken", *refreshToken, 172800, "", "localhost", false, false)

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
