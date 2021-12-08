package response

import "time"

type LoginResponse struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Username  string
	RoleId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
