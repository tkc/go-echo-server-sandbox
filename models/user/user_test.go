package userModel

import (
		"testing"
		"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
		//u := User{}
		//u.Migrate()
}

func TestCreate(t *testing.T) {
		u := User{}
		user := u.Create("test", 30)
		assert.Equal(t, user.Name, "test")
		assert.Equal(t, user.Age, 30)
}

func TestFetch(t *testing.T) {
		//u := User{}
		//user := u.Fetch(10);
		//t.Log(user)
		//assert.Equal(t, user.Id, 2)//TODO Not equal
}

func TestAll(t *testing.T) {
		u := User{}
		users := u.All();
		t.Log(users)
		for i, user := range users {
				t.Log(i)
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
		users := u.FetchByName("test");
		for _, v := range users {
				assert.Equal(t, v.Name, "test")
		}
}

func TestFetchByNameAge(t *testing.T) {
		u := User{}
		users := u.FetchByNameAndAge("test", 30);
		for _, v := range users {
				assert.Equal(t, v.Name, "test")
				assert.Equal(t, v.Age, 30)
		}
}

func TestMapByName(t *testing.T) {
		u := User{}
		users := u.MapByName("test");
		for _, v := range users {
				assert.Equal(t, v.Name, "test")
		}
}

func TestUpdateName(t *testing.T) {
		u := User{}
		user := u.UpdateName(4, "Hoo");
		assert.Equal(t, user.Name, "Hoo")
}

func TestDelete(t *testing.T) {
		//u := User{}
		//u.Delete(2);
}
