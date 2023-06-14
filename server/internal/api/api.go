package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type API struct {
	gin  *gin.Engine
	port uint
}

func NewAPI(port uint) *API {
	api := &API{
		gin:  gin.Default(),
		port: port,
	}

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
