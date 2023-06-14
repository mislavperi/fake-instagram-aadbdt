package bootstrap

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/cmd/api/config"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql"
)

func newPostgresRepository() (*psql.Repository, error) {
	repository, err := psql.NewRepository(config.Cfg.Database.Host, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Name, config.Cfg.Database.Port)
	if err != nil {
		return nil, err
	}
	return repository, err
}

func Api() (*api.API, error) {
	_, err := newPostgresRepository()
	if err != nil {
		panic(err)
	}
	api := api.NewAPI(8080)
	return api, nil
}
