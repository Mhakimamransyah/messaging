package users_test

import (
	"errors"
	"messaging/business/users"
	"messaging/business/users/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	ID       = 1
	Name     = "M.Hakim Amransyah"
	Username = "mhakim"
	Phone    = "081271286874"
)

var (
	usersData    users.Users
	usersSpec    users.UsersSpec
	listUserData []*users.Users

	user_repo    mocks.Repository
	user_service users.Services
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetUser(t *testing.T) {
	t.Run("Expects succes get user by username", func(t *testing.T) {
		user_repo.On("Get", mock.AnythingOfType("string")).Return(&usersData, nil).Once()
		res, err := user_service.GetUser(Username)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})

	t.Run("Expects use not found and get error", func(t *testing.T) {
		user_repo.On("Get", Username).Return(nil, errors.New("Something error in repos")).Once()
		user, err := user_service.GetUser(Username)
		assert.NotNil(t, err)
		assert.Nil(t, user)
	})
}

func TestRegistersNewUser(t *testing.T) {

	t.Run("Expects Invalid Specs", func(t *testing.T) {
		failed_specs := users.UsersSpec{
			Name:  "Hakim",
			Phone: "082170129870",
		}
		err := user_service.RegistersNewUser(&failed_specs)
		assert.NotNil(t, err)
	})

	t.Run("Expects something wrong when insert to database", func(t *testing.T) {
		user_repo.On("CreateUser", mock.AnythingOfType("*users.Users")).Return(errors.New("Error")).Once()
		err := user_service.RegistersNewUser(&usersSpec)
		assert.NotNil(t, err)
	})

	t.Run("Expects success create new users", func(t *testing.T) {
		user_repo.On("CreateUser", mock.AnythingOfType("*users.Users")).Return(nil).Once()
		err := user_service.RegistersNewUser(&usersSpec)
		assert.Nil(t, err)
	})
}

func TestGetAllUser(t *testing.T) {
	t.Run("Expects Error while reading all users data", func(t *testing.T) {
		user_repo.On("GetAll").Return(nil, errors.New("Error")).Once()
		res, err := user_service.GetAllUser()
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})
	t.Run("Expects Success while reading all users data", func(t *testing.T) {
		user_repo.On("GetAll").Return(listUserData, nil).Once()
		res, err := user_service.GetAllUser()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func setup() {
	usersData = users.Users{
		ID:       ID,
		Name:     Name,
		Username: Username,
		Phone:    Phone,
	}

	usersSpec = users.UsersSpec{
		Name:     Name,
		Username: Username,
		Phone:    Phone,
	}

	listUserData = append(listUserData, &usersData)
	user_service = users.InitUserService(&user_repo)
}
