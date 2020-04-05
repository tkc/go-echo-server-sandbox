package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tkc/go-echo-server-sandbox/config"
)

func Connect() *gorm.DB {

	connection := config.GetDataBaseAccess()
	db, err := gorm.Open(
		"mysql",
		connection,
	)
	if err != nil {
		log.Print(connection)
		log.Print(err)
		panic("failed to connect database")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	return db
}
