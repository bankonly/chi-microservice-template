package configs

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ProdLabel string `yaml:"prodLabel"`
	DevLabel  string `yaml:"devLabel"`
}

var AppConf AppConfig

func LoadAppConfig() {
	dest := "./internal/configs/yaml/" + Environment.MODE + ".yaml"
	viper.SetConfigFile(dest)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&AppConf); err != nil {
		log.Fatal(err)
	}
}
