package main

import (
	"github.com/tkc/go-echo-server-sandbox/db"
	"github.com/tkc/go-echo-server-sandbox/models/user"
)

func main() {
	db := db.Connect()
	db.AutoMigrate(&userModel.User{})
}
