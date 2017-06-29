package userModel

import (
	"github.com/tkc/go-echo-server-sandbox/db"
)

var user User
var users [] User

func (u *User) GetName() string {
	return u.Name
}

func (u *User) All(limit int, offset int) []User {
	db := db.Connect()
	db.Limit(limit).Offset(offset).Order("id desc,name, age").Find(&users)
	return users
}

func (u *User) Create(name string, age int) User {
	db := db.Connect()
	user = User{Name: name, Age: age}
	db.Create(&user)
	return user
}

func (u *User) UpdateName(id int64, name string) User {
	db := db.Connect()
	user = User{Id: id}
	db.Model(&user).Update("name", name)
	return user
}

func (u *User) Fetch(id int64) User {
	db := db.Connect()
	user = User{Id: id}
	db.First(&user)
	return user
}

func (u *User) FetchByName(name string) []User {
	db := db.Connect()
	db.Where("name = ?", name).First(&users)
	return users
}

func (u *User) FetchByLikeName(name string) []User {
	db := db.Connect()
	db.Where("name LIKE ?", "%"+name+"% ").Find(&users)
	return users
}

func (u *User) FetchByNameAndAge(name string, age int) []User {
	db := db.Connect()
	db.Where(&User{Name: name, Age: age}).First(&users)
	return users
}

func (u *User) MapByName(name string) []User {
	db := db.Connect()
	db.Where(map[string]interface{}{"name": name}).Find(&users)
	return users
}

func (u *User) Delete(id int64) {
	db := db.Connect()
	user = User{Id: id}
	db.Delete(&user)
}
