package db

import (
	"github.com/tkc/go-echo-server-sandbox/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	connection := env.GetDataBaseAccess()
	db, err := gorm.Open(
		"mysql",
		connection,
	)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
