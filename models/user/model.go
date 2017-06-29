package userModel

import (
	"time"
)

type User struct {
	Id        int64
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

