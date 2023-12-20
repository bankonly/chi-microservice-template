package main

import (
	"ecm-api-template/internal/configs"
	"ecm-api-template/pkg/storages"
	"log"
	"os"
)

// Danger!!
// MAKE SURE YOU BACKUP FOR YOUR DATABASE
// For development purpose only

func main() {
	configs.LoadEnvironmentConf()
	storages.NewPostgres()
	db := storages.GetPostgresDB()

	sqlFile, err := os.ReadFile("./cmd/migration/migration.psql")
	if err != nil {
		log.Fatal(err)
	}
	execute := db.Exec(string(sqlFile))
	if execute.Error != nil {
		log.Fatal(err)
	}
	log.Println("Migration completed")
}
