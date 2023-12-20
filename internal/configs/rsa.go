package configs

import (
	"log"
	"os"

	"github.com/bankonly/go-pkg/v1/encryption"
)

func LoadRSAConf() {
	pkFile, err := os.ReadFile(Environment.SELF_PRIVATE_KEY_PATH)
	if err != nil {
		log.Fatal(err)
	}

	pbkFile, err := os.ReadFile(Environment.SELF_PUBLIC_KEY_PATH)
	if err != nil {
		log.Fatal(err)
	}

	encryption.SetRSAKey(string(pkFile), string(pbkFile))
}
