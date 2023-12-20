package main

import (
	"ecm-api-template/internal/configs"
	"log"

	"github.com/bankonly/go-pkg/v1/encryption"
)

func main() {
	configs.LoadEnvironmentConf()

	err := encryption.NewRSA(encryption.RSAConfig{
		Filename:        configs.Environment.KEY_NAME,
		DestinationPath: "./secret/",
		BackupPath:      "./cmd/key-pair/backup/",
	})
	if err != nil {
		panic(err)
	}

	configs.LoadRSAConf()

	encData, err := encryption.RSAEncAESRandomKey("f933d0b8-8002-4623-8330-62943d95b359")
	if err != nil {
		panic(err)
	}

	log.Println("IV =", encData.EncryptedData.IvInfo.IvString)
	log.Println("ENK-SESSION =", encData.EncryptedData.Data)
	log.Println("ENK =", encData.EncryptKey)

	log.Println("Key-Pair Generated")
}
