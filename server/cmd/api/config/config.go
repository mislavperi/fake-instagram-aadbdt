package config

import (
	"github.com/kelseyhightower/envconfig"
)

var Cfg Config

type Config struct {
	Database DatabaseConfig `split_words:"true"`
	Aws      AwsConfig      `split_words:"true"`
	Github   GithubConfig   `split_words:"true"`
	Auth     AuthConfig     `split_words:"true"`
}

type DatabaseConfig struct {
	Host     string `split_words:"true"`
	User     string `split_words:"true"`
	Password string `split_words:"true"`
	Name     string `split_words:"true"`
	Port     string `split_words:"true"`
}

type AwsConfig struct {
	AccessKeyId     string `split_words:"true"`
	SecretAccessKey string `split_words:"true"`
	Bucket          string `split_words:"true"`
	Region          string `split_words:"true"`
}

type GithubConfig struct {
	ClientID     string `split_words:"true"`
	ClientSecret string `split_words:"true"`
}

type AuthConfig struct {
	SecretKey string `split_words:"true"`
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
