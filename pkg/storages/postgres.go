package storages

import (
	"database/sql"
	"ecm-api-template/internal/configs"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetPostgresDB() *gorm.DB {
	if db == nil {
		log.Fatal("database_not_initialized")
	}

	return db
}

func BeginTxnReadUncommitted() *gorm.DB {
	if db == nil {
		log.Fatal("database_not_initialized")
	}

	return db.Begin(&sql.TxOptions{Isolation: sql.LevelReadUncommitted})
}

func NewPostgres() {
	var err error
	db, err = gorm.Open(postgres.Open(configs.Environment.PG_URI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB is connected", strings.Split(configs.Environment.PG_URI, "@")[1])
}
