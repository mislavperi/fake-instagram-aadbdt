package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	customerrors "github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
)

//go:generate mockery --output=./tests/mocks --name=UserService
type UserService interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	Login(username string, password string) (*string, *string, error)
	GetUserInformation(id int) (*models.User, error)
	SelectUserPlan(id int, plan models.Plan) error
	AuthenticateGoogleUser(credentials models.GoogleToken) (*string, *string, error)
	AuthenticateGithubUser(credentials models.GHCredentials) (*string, *string, error)
	GetAllUsers(adminID int) ([]models.User, error)
	InsertAdminPlanChange(adminID int, userID int, planID int) error
	GetUserLogs(userID int, adminID int) ([]models.Log, error)
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
		identifier, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user, err := c.UserService.GetUserInformation(identifier)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
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

		identifier, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = c.UserService.SelectUserPlan(identifier, plan)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *UserController) AdminUserPlanChange() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Query("id")
		planID := ctx.Query("planId")

		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		planIDInt, err := strconv.Atoi(planID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		adminID, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = c.UserService.InsertAdminPlanChange(adminID, userIDInt, planIDInt)
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
		accessToken, refreshToken, err := c.UserService.Login(user.Username, user.Password)
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

func (c *UserController) GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identifier, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		users, err := c.UserService.GetAllUsers(identifier)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}

func (c *UserController) GetUserLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userID int
		incReq := ctx.Query("id")
		err := json.Unmarshal([]byte(incReq), &userID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		identifier, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		logs, err := c.UserService.GetUserLogs(userID, identifier)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, logs)
	}
}
