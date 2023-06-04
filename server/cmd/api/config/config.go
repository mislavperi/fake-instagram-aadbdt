package config

import (
	"github.com/kelseyhightower/envconfig"
)

var Cfg Config

type Config struct {
	Database DatabaseConfig `split_words:"true"`
}

type DatabaseConfig struct {
	Host         string `split_words:"true"`
	User         string `split_words:"true"`
	Password     string `split_words:"true"`
	DatabaseName string `split_words:"true"`
	Port         string `split_words:"true"`
}

func DummyLoad() error {
	err := Load(&Cfg)
	return err
}

func Load(cfg interface{}) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}
	return nil
}
