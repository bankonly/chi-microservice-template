package main

import (
	"ecm-api-template/internal"
	"ecm-api-template/internal/configs"
	"ecm-api-template/pkg/storages"

	"github.com/bankonly/go-pkg/v1/validator"
)

func main() {
	configs.LoadEnvironmentConf()
	configs.LoadRSAConf()
	configs.LoadRedisConf()
	configs.LoadAppConfig()

	storages.NewRedis()
	storages.NewPostgres()
	validator.New(validator.ValidatorConfig{})
	internal.NewServer()
}
