package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Environments struct {
	MODE                  string `env:"MODE,required"`
	PORT                  string `env:"PORT,required"`
	PG_URI                string `env:"PG_URI,required"`
	REDIS_URI             string `env:"REDIS_URI,required"`
	KEY_NAME              string `env:"KEY_NAME,required"`
	SELF_PRIVATE_KEY_PATH string `env:"SELF_PRIVATE_KEY_PATH,required"`
	SELF_PUBLIC_KEY_PATH  string `env:"SELF_PUBLIC_KEY_PATH,required"`
}

var Environment Environments

func LoadEnvironmentConf() {
	if err := env.Parse(&Environment); err != nil {
		log.Fatal(err.Error())
	}
}
