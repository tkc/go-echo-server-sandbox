package userModel

import (
	"github.com/tkc/go-echo-server-sandbox/db"
)

func (u *User) GetName() string {
	return u.Name
}

func (u *User) All(limit int, offset int) []User {
	var users [] User
	db := db.Connect()
	db.Limit(limit).Offset(offset).Order("id desc").Find(&users)
	return users
}

func (u *User) Create(name string, age int) User {
	var user User
	db := db.Connect()
	user = User{Name: name, Age: age}
	db.Create(&user)
	return user
}

func (u *User) Update(id int64, name string,age int) User {
	var user User
	db := db.Connect()
	user = User{Id: id}
	db.Model(&user).Update(
		User{
			Name:name,
			Age:age,
		})
	return user
}

func (u *User) Fetch(id int64) User {
	var user User
	db := db.Connect()
	user = User{Id: id}
	db.First(&user)
	return user
}

func (u *User) FetchByName(name string) []User {
	var users [] User
	db := db.Connect()
	db.Where("name = ?", name).First(&users)
	return users
}

func (u *User) FetchByLikeName(name string) []User {
	var users [] User
	db := db.Connect()
	db.Where("name LIKE ?", "%"+name+"% ").Find(&users)
	return users
}

func (u *User) FetchByNameAndAge(name string, age int) []User {
	var users [] User
	db := db.Connect()
	db.Where(&User{Name: name, Age: age}).First(&users)
	return users
}

func (u *User) MapByName(name string) []User {
	var users [] User
	db := db.Connect()
	db.Where(map[string]interface{}{"name": name}).Find(&users)
	return users
}

func (u *User) Delete(id int64) {
	if id >0 {
		var user User
		db := db.Connect()
		user = User{Id: id}
		db.Delete(&user)
	}
}
