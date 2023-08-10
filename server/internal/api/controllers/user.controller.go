package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	customerrors "github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
)

const (
	client_id     = "f4a7a59f8e527f183bcf"
	client_secret = "7d5ec2c82ebfd99f412ea4a71addb2e45b98d494"
)

type UserService interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	Login(email string, password string) (*string, *string, error)
	GetUserInformation(email string) (*models.User, error)
	SelectUserPlan(email string, plan models.Plan) error
	AuthenticateGoogleUser(credentials models.GoogleToken) (*string, *string, error)
	AuthenticateGithubUser(credentials models.GHCredentials) (*string, *string, error)
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
		user, err := c.UserService.GetUserInformation(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *UserController) SetUserPlan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var plan models.Plan
		if err := ctx.BindJSON(&plan); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err := c.UserService.SelectUserPlan(ctx.GetHeader("Identifier"), plan)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *UserController) RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser models.User
		if err := ctx.BindJSON(&newUser); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
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
		accessToken, refreshToken, err := c.UserService.Login(user.Email, user.Password)
		if err != nil {
			if customerrors.IsInvalidCredentialsError(err) {
				ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid credentials"))
				return
			}
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.SetCookie("accessToken", *accessToken, 3600, "", "localhost", false, false)
		ctx.SetCookie("refreshToken", *refreshToken, 172800, "", "localhost", false, false)

		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *UserController) LoginGithub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ghCredentials models.GHCredentials
		ctx.BindJSON(&ghCredentials)
		accessToken, refreshToken, err := c.UserService.AuthenticateGithubUser(ghCredentials)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.SetCookie("accessToken", *accessToken, 3600, "/", "localhost", false, false)
		ctx.SetCookie("refreshToken", *refreshToken, 172800, "/", "localhost", false, false)

		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *UserController) LoginGoogle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var googleCreds models.GoogleToken

		accessToken, refreshToken, err := c.UserService.AuthenticateGoogleUser(googleCreds)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.SetCookie("accessToken", *accessToken, 3600, "/", "localhost", false, true)
		ctx.SetCookie("refreshToken", *refreshToken, 172800, "/", "localhost", false, true)

		ctx.JSON(http.StatusOK, nil)
	}
}
