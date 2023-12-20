package configs

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type RedisKeyConf struct {
	Key       string        `yaml:"key"`
	ExpireMin time.Duration `yaml:"expireMin"`
}

type RedisConfig struct {
	SessionInfo   RedisKeyConf `yaml:"sessionInfo"`
	SessionVector RedisKeyConf `yaml:"sessionVector"`
}

var RedisConf RedisConfig

func LoadRedisConf() {
	dest := "./internal/configs/yaml/redis.yaml"
	viper.SetConfigFile(dest)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&RedisConf); err != nil {
		log.Fatal(err)
	}
}
