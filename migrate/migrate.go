package main

import (
	"github.com/tkc/go-echo-server-sandbox/models/user"
	"github.com/tkc/go-echo-server-sandbox/db"
)

func main() {
	db := db.Connect()
	db.AutoMigrate(&userModel.User{})
}
