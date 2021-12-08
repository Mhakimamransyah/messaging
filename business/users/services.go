package users

import (
	"errors"
	"messaging/business"
	"messaging/util/hashing"
	"messaging/util/validator"
)

type UserService struct {
	User_repo Repository
}

func InitUserService(repository Repository) *UserService {
	return &UserService{
		User_repo: repository,
	}
}

type UsersSpec struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Phone    string `form:"phone" validate:"required"`
	RoleId   int    `form:"role_id" validate:"required"`
}

func (service *UserService) RegistersNewUser(users *UsersSpec) error {
	err := validator.GetValidator().Struct(users)
	if err != nil {
		return business.ErrInvalidSpec
	}
	err = service.User_repo.CreateUser(NewUser(users))
	if err != nil {
		return errors.New("Register Failed")
	}
	return nil
}

func (service *UserService) Login(username, password string) (*Users, error) {
	user, err := service.User_repo.Login(username, password)
	if err != nil {
		return nil, err
	}
	if hashing.CompareHash(password, user.Password) {
		return user, nil
	} else {
		return nil, business.ErrLogin
	}
}

func (service *UserService) GetUser(username string) (*Users, error) {
	return nil, nil
}

func (service *UserService) RemoveUser(username string) error {
	return nil
}

func (service *UserService) ModifyUser(user *UsersSpec) error {
	return nil
}
