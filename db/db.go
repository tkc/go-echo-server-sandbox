package db

import (
		"github.com/tkc/go-echo-server-sandbox/config"
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
		connection := config.DbUser + ":" + config.DbPass + "@tcp([" + config.DbHost + "]:" + config.DbPort + ")/" + config.DbName + "?charset=utf8&parseTime=True&loc=" + config.DbLocate
		db, err := gorm.Open(
				"mysql",
				connection,
		)
		if err != nil {
				panic("failed to connect database")
		}
		return db
}
