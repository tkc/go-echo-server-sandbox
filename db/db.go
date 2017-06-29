package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-echo-server-sandbox/config"
)

func Connect() *gorm.DB {
	connection := config.GetDataBaseAccess()
	db, err := gorm.Open(
		"mysql",
		connection,
	)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
