package api

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/cmd/api/config"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql"
)

func newPostgresRepository() (*psql.Repository, error) {
	repository, err := psql.NewRepository(config.Cfg.Database.Host, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.DatabaseName, config.Cfg.Database.Port)
	if err != nil {
		return nil, err
	}
	return repository, err
}

func Api() {
	repository, err := newPostgresRepository()
	if err != nil {
		panic(err)
	}
}
