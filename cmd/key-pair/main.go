package main

import (
	"ecm-api-template/internal/configs"
	"log"

	"github.com/bankonly/go-pkg/v1/encryption"
)

func main() {
	configs.LoadEnvironmentConf()
	configs.LoadRSAConf()

	err := encryption.NewRSA(encryption.RSAConfig{
		Filename:        configs.Environment.KEY_NAME,
		DestinationPath: "./secret/",
		BackupPath:      "./cmd/key-pair/backup/",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Key-Pair Generated")
}
