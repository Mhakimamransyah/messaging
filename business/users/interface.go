package users

type Services interface {
	RegistersNewUser(users *UsersSpec) error
	Login(username, password string) (*Users, error)
	GetUser(username string) (*Users, error)
	RemoveUser(username string) error
	ModifyUser(user *UsersSpec) error
}

type Repository interface {
	CreateUser(user *Users) error
	Login(username, password string) (*Users, error)
	Get(username string) (*Users, error)
	Update(user *Users) error
	Delete(user *Users) error
}
