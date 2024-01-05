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
		FieldNullable:    true,
		FieldWithTypeTag: true,
		ModelPkgPath:     "internal/models/entity",
	})

	g.WithOpts(gen.FieldType("tags", "*[]string"))
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
