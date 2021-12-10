package users

import (
	"log"
	"messaging/business"
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
	Name     string `form:"name" json:"name" validate:"required,max=20"`
	Username string `form:"username" json:"username" validate:"required,max=10"`
	Phone    string `form:"phone" json:"phone" validate:"max=20"`
}

func (service *UserService) GetAllUser() ([]*Users, error) {
	res, err := service.User_repo.GetAll()
	if err != nil {
		return nil, business.ErrInternalServerError
	}
	return res, nil
}

func (service *UserService) RegistersNewUser(users *UsersSpec) error {
	err := validator.GetValidator().Struct(users)
	if err != nil {
		log.Printf("%s", err.Error())
		return business.ErrInvalidSpec
	}
	err = service.User_repo.CreateUser(NewUser(users))
	if err != nil {
		log.Printf("%s", err.Error())
		return business.ErrInvalidRequest
	}
	return nil
}

func (service *UserService) GetUser(username string) (*Users, error) {
	res, err := service.User_repo.Get(username)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, business.ErrInvalidRequest
	}
	return res, nil
}
