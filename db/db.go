package db

import (
	"github.com/jinzhu/gorm"
	"github.com/tkc/go-echo-server-sandbox/config"
	_ "github.com/go-sql-driver/mysql"
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

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	return db
}
