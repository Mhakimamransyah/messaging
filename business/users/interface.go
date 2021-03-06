package users

type Services interface {
	RegistersNewUser(users *UsersSpec) error
	GetUser(username string) (*Users, error)
	GetAllUser() ([]*Users, error)
}

type Repository interface {
	CreateUser(user *Users) error
	Login(username, password string) (*Users, error)
	Get(username string) (*Users, error)
	GetUserById(id_user int) (*Users, error)
	GetAll() ([]*Users, error)
	Update(user *Users) error
	Delete(user *Users) error
}
