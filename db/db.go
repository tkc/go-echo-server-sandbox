package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tkc/go-echo-server-sandbox/config"
)

var dbConnection *gorm.DB

func Connect() *gorm.DB {

	return dbConnection

	if dbConnection != nil {
		return dbConnection
	} else {

		connection := config.GetDataBaseAccess()
		db, err := gorm.Open(
			"mysql",
			connection,
		)
		if err != nil {
			panic("failed to connect database")
		}

		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		dbConnection = db
		return dbConnection
	}
}
