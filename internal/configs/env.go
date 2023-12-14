package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Environments struct {
	PG_URI           string `env:"PG_URI,required"`
	REDIS_URI        string `env:"REDIS_URI,required"`
	SELF_PRIVATE_KEY string `env:"SELF_PRIVATE_KEY,required"`
	SELF_PUBLIC_KEY  string `env:"SELF_PUBLIC_KEY,required"`
}

var Environment Environments

func _LoadEnvironmentConf() {
	if err := env.Parse(&Environment); err != nil {
		log.Fatal(err.Error())
	}
}
