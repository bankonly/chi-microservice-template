package envConf

import (
	"log"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Environment struct {
	PG_URI    string `env:"PG_URI,required"`
	REDIS_URI string `env:"REDIS_URI,required"`
}

var ValueOf Environment

func Load() {
	if err := env.Parse(&ValueOf); err != nil {
		log.Fatal(err.Error())
	}
}
