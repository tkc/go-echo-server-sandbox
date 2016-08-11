package db

import (
		"github.com/tkc/go-echo-server-sandbox/config"
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
		db, err := gorm.Open(
				"mysql",
				config.DbUser + ":" + config.DbPass + "@/" + config.DbName + "?charset=utf8&parseTime=True&loc=" + config.DbHost,
		)
		if err != nil {
				panic("failed to connect database")
		}
		return db
}
