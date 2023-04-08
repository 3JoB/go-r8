package main

import (
	"github.com/3JoB/go-r8/internal/db"
	"gorm.io/gen"
)

type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main(){
	gg := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	})

	gg.UseDB(db.NewDB())
}