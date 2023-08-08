package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	customerrors "github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
)

const (
	client_id     = "f4a7a59f8e527f183bcf"
	client_secret = ""
)

type UserService interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	Login(email string, password string) (*string, *string, error)
	GetUserInformation(email string) (*models.User, error)
	SelectUserPlan(email string, plan models.Plan) error
	AuthenticateGoogleUser(user models.GoogleUser) (*string, *string, error)
	AuthenticateGithubUser(user models.GHUser) (*string, *string, error)
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
		accessToken, refreshToken, err := c.UserService.AuthenticateGithubUser(ghUser)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.SetCookie("accessToken", *accessToken, 3600, "/", "localhost", false, true)
		ctx.SetCookie("refreshToken", *refreshToken, 172800, "/", "localhost", false, true)

		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *UserController) LoginGoogle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var incCreds models.GoogleToken
		var googleUser models.GoogleUser

		err := ctx.BindJSON(&incCreds)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return

		}
		token, err := jwt.ParseWithClaims(
			incCreds.GoogleJWT,
			&googleUser,
			func(t *jwt.Token) (interface{}, error) {
				resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
				if err != nil {
					return nil, err
				}
				dat, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				}

				cemKey := map[string]string{}
				err = json.Unmarshal(dat, &cemKey)
				if err != nil {
					return nil, err
				}
				pem, ok := cemKey[fmt.Sprintf("%s", t.Header["kid"])]
				if !ok {
					return nil, err
				}
				key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
				if err != nil {
					return nil, err
				}
				return key, nil
			},
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		claims, ok := token.Claims.(*models.GoogleUser)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		accessToken, refreshToken, err := c.UserService.AuthenticateGoogleUser(*claims)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.SetCookie("accessToken", *accessToken, 3600, "/", "localhost", false, true)
		ctx.SetCookie("refreshToken", *refreshToken, 172800, "/", "localhost", false, true)

		ctx.JSON(http.StatusOK, nil)
	}
}
