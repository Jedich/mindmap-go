package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"mindmap-go/app/models"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:11@(127.0.0.1:6603)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `models.User` following conventions
	g.ApplyBasic(models.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `models.User` and `models.Company`
	g.ApplyInterface(func(Querier) {}, models.User{})

	// Generate the code
	g.Execute()
}
