package users

import (
	"messaging/util/hashing"
	"time"
)

type Users struct {
	ID        int
	Name      string
	Email     string
	Username  string
	Password  string
	Phone     string
	RolesId   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUser(users *UsersSpec) *Users {
	hashPassword, _ := hashing.Generate(users.Password)
	return &Users{
		Name:      users.Name,
		Email:     users.Email,
		Username:  users.Username,
		Password:  hashPassword,
		Phone:     users.Phone,
		RolesId:   users.RoleId,
		CreatedAt: time.Now(),
	}
}
