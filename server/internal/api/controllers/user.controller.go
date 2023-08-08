package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	customerrors "github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
)

const (
	client_id     = "fd9e7705baa36a1a120d"
	client_secret = "f33505f8e0311cd11c651a71b9cd2ddb04f1edfe"
)

type UserService interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	Login(username string, password string) (*string, *string, error)
	GetUserInformation(username string) (*models.User, error)
	SelectUserPlan(username string, plan models.Plan) error
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
		err := c.UserService.SelectUserPlan("", plan)
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

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *UserController) LoginGithub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ghCredentials models.GHCredentials
		var ghToken models.GHToken
		var ghUser models.GHUser
		ctx.BindJSON(&ghCredentials)

		body := models.GHCredsReq{
			Code:         ghCredentials.Code,
			ClientID:     client_id,
			ClientSecret: client_secret,
		}
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, err := http.NewRequest(
			"POST",
			"https://github.com/login/oauth/access_token",
			bytes.NewBuffer(bodyJSON),
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = json.NewDecoder(resp.Body).Decode(&ghToken)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		infoRequest, err := http.NewRequest("GET", "https://api.github.com/user", nil)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		infoRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ghToken.AccessToken))
		infoResp, err := http.DefaultClient.Do(infoRequest)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = json.NewDecoder(infoResp.Body).Decode(&ghUser)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(200, ghUser)
	}
}
