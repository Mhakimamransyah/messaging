package users

import (
	"time"
)

type Users struct {
	ID        int
	Name      string
	Username  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUser(users *UsersSpec) *Users {
	return &Users{
		Name:      users.Name,
		Username:  users.Username,
		Phone:     users.Phone,
		CreatedAt: time.Now(),
	}
}
