package userModel

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	u := User{}
	user := u.Create("test", 30)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Age, 30)
}

func TestFetch(t *testing.T) {
	u := User{}
	user := u.Fetch(10)
	assert.Equal(t, int(user.Id), 10)
}

func TestAll(t *testing.T) {
	u := User{}
	users := u.All(10, 0)
	for i, user := range users {
		t.Log(i)
		t.Log(user.Name)
		t.Log(user)
	}
}

func TestGetName(t *testing.T) {
	u := User{}
	u.Name = "test"
	name := u.GetName()
	assert.Equal(t, name, "test")
}

func TestFetchByName(t *testing.T) {
	u := User{}
	users := u.FetchByName("test")
	for _, v := range users {
		assert.Equal(t, v.Name, "test")
	}
}

func TestFetchByNameAge(t *testing.T) {
	u := User{}
	users := u.FetchByNameAndAge("test", 30)
	for _, v := range users {
		assert.Equal(t, v.Name, "test")
		assert.Equal(t, v.Age, 30)
	}
}

func TestMapByName(t *testing.T) {
	u := User{}
	users := u.MapByName("test")
	for _, v := range users {
		assert.Equal(t, v.Name, "test")
	}
}

func TestUpdateName(t *testing.T) {
	u := User{}
	user := u.UpdateName(4, "Foo")
	assert.Equal(t, user.Name, "Foo")
}
