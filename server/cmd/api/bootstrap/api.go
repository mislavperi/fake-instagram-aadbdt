package bootstrap

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/cmd/api/config"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/repositories"
)

func Api() (*api.API, error) {
	db, err := psql.NewDatabaseConnection(config.Cfg.Database.Host, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Name, config.Cfg.Database.Port)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	if err != nil {
		panic(err)
	}
	api := api.NewAPI(userController, 8080)
	return api, nil
}
