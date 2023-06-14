package main

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/cmd/api/bootstrap"
	"github.com/mislavperi/fake-instagram-aadbdt/server/cmd/api/config"
	"github.com/mislavperi/fake-instagram-aadbdt/server/utils"
)

func main() {
	err := config.DummyLoad()
	if err != nil {
		panic(err)
	}

	api, err := bootstrap.Api()

	utils.Run(
		api,
	)
}
