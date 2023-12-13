package main

import (
	"log"

	"github.com/bankonly/go-pkg/v1/encryption"
)

func main() {
	err := encryption.NewRSA(encryption.RSAConfig{
		Filename:        "self_key",
		DestinationPath: "./",
		BackupPath:      "./cmd/key-pair/backup/",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Key-Pair Generated")
}
