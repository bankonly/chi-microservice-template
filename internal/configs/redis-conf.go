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
	SessionInfo struct {
		Key       string        `yaml:"key"`
		ExpireMin time.Duration `yaml:"expireMin"`
	} `yaml:"session_info"`
}

var RedisConf RedisConfig

func _LoadRedisConf() {
	dest := "./internal/configs/yaml/redis.yaml"
	viper.SetConfigFile(dest)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&RedisConf); err != nil {
		log.Fatal(err)
	}
}
