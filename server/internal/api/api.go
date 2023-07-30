package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers"
)

type API struct {
	gin  *gin.Engine
	port uint
}

func NewAPI(userController *controllers.UserController, port uint) *API {
	api := &API{
		gin:  gin.Default(),
		port: port,
	}
	api.registerRoutes(userController)
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

func (a *API) registerRoutes(userController *controllers.UserController) {
	userGroup := a.gin.Group("/user")
	userGroup.POST("/register", userController.RegisterUser())
	userGroup.POST("/login", userController.Login())

}
