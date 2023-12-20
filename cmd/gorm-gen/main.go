package main

import (
	"ecm-api-template/internal/configs"
	"ecm-api-template/pkg/storages"

	"gorm.io/gen"
)

func main() {
	configs.LoadEnvironmentConf()
	storages.NewPostgres()
	db := storages.GetPostgresDB()

	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: "internal/models/entity",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
