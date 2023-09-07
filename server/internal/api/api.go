package api

import (
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type API struct {
	Gin  *gin.Engine
	port uint
}

func NewAPI(userController *controllers.UserController, planController *controllers.PlanController, pictureController *controllers.PictureController, uploadController *controllers.UploadController, port uint) *API {
	api := &API{
		Gin:  gin.Default(),
		port: port,
	}
	api.Gin.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "http://localhost", "http://localhost:5173", "http://localhost:9090", "http://localhost:9091"},
				AllowMethods:     []string{"POST", "GET"},
				AllowHeaders:     []string{"Content-Type", "Accept", "Authorization", "Refresh"},
				AllowCredentials: true,
			},
		),
	)
	api.registerRoutes(userController, planController, pictureController, uploadController)
	return api
}

func (a *API) Start(ctx context.Context) {
	errs := make(chan error, 1)

	go func() {
		errs <- a.Gin.Run(fmt.Sprintf(":%d", a.port))
	}()

	select {
	case <-errs:
		return
	case <-ctx.Done():
		return
	}
}

func (a *API) registerRoutes(userController *controllers.UserController, planController *controllers.PlanController, pictureController *controllers.PictureController, uploadController *controllers.UploadController) {
	metricsGroup := a.Gin.Group("/metrics")
	metricsGroup.GET("/", func(ctx *gin.Context) {
		h := promhttp.Handler()

		h.ServeHTTP(ctx.Writer, ctx.Request)
	})

	authGroup := a.Gin.Group("/api/auth")
	authGroup.POST("/register", userController.RegisterUser())
	authGroup.POST("/login", userController.Login())
	authGroup.POST("/gh_login", userController.LoginGithub())
	authGroup.POST("/g_login", userController.LoginGoogle())

	userGroup := a.Gin.Group("/api/user")
	userGroup.Use(middlewares.JwtTokenCheck())
	userGroup.GET("/whoami", userController.Whoami())
	userGroup.POST("/select", userController.SetUserPlan())

	adminGroup := a.Gin.Group("/api/admin")
	adminGroup.Use(middlewares.JwtTokenCheck())
	adminGroup.GET("/users", userController.GetAllUsers())
	adminGroup.GET("/statistics", uploadController.GetUserStatistics())
	adminGroup.GET("/changePlan", userController.AdminUserPlanChange())
	adminGroup.GET("/userPictures", pictureController.GetSpecificUserImages())
	adminGroup.GET("/userlogs", userController.GetUserLogs())

	planGroup := a.Gin.Group("/api/plans")
	planGroup.GET("/get", planController.GetPlans())

	unauthorizedPictureGroup := a.Gin.Group("/api/public/picture")
	unauthorizedPictureGroup.GET("/get", pictureController.GetImages())

	pictureGroup := a.Gin.Group("/api/picture")
	pictureGroup.Use(middlewares.JwtTokenCheck())
	pictureGroup.POST("/upload", pictureController.UploadImage())
	pictureGroup.GET("/userImages", pictureController.GetUserImages())
	pictureGroup.GET("/info", pictureController.GetPictureByID())
	pictureGroup.POST("/update", pictureController.UpdateImage())
	pictureGroup.POST("/edited", pictureController.GetEditedImage())

	uploadGroup := a.Gin.Group("/api/statistics")
	uploadGroup.Use(middlewares.JwtTokenCheck())
	uploadGroup.GET("/get", uploadController.GetStatistics())
}
