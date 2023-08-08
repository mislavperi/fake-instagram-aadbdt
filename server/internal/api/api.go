package api

import (
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/middlewares"
)

type API struct {
	gin  *gin.Engine
	port uint
}

func NewAPI(userController *controllers.UserController, planController *controllers.PlanController, port uint) *API {
	api := &API{
		gin:  gin.Default(),
		port: port,
	}
	api.gin.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
				AllowMethods:     []string{"POST", "GET"},
				AllowHeaders:     []string{"Content-Type", "Accept", "Authorization", "Refresh"},
				AllowCredentials: true,
			},
		),
	)
	api.registerRoutes(userController, planController)
	return api
}

func (a *API) Start(ctx context.Context) {
	errs := make(chan error, 1)

	go func() {
		errs <- a.gin.Run(fmt.Sprintf(":%d", a.port))
	}()

	select {
	case <-errs:
		return
	case <-ctx.Done():
		return
	}
}

func (a *API) registerRoutes(userController *controllers.UserController, planController *controllers.PlanController) {
	authGroup := a.gin.Group("/auth")
	authGroup.POST("/register", userController.RegisterUser())
	authGroup.POST("/login", userController.Login())
	authGroup.POST("/gh_login", userController.LoginGithub())
	authGroup.POST("/g_login", userController.LoginGoogle())

	userGroup := a.gin.Group("/user")
	userGroup.Use(middlewares.JwtTokenCheck())
	userGroup.GET("/whoami", userController.Whoami())
	userGroup.POST("/select", userController.SetUserPlan())

	planGroup := a.gin.Group("/plans")
	planGroup.GET("/get", planController.GetPlans())
}
