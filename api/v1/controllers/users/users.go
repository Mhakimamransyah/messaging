package controllers

import (
	"errors"
	"log"
	"messaging/api/common"
	"messaging/business/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	User_service users.Services
}

func InitUserController(service users.Services) *UsersController {
	return &UsersController{
		User_service: service,
	}
}

func (controller *UsersController) Register(c echo.Context) error {
	usersSpec := users.UsersSpec{}
	c.Bind(&usersSpec)
	err := controller.User_service.RegistersNewUser(&usersSpec)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessCreated())
}

func (controller *UsersController) GetUsersByUsername(c echo.Context) error {
	username := c.Param("username")
	if username != "" {
		res, err := controller.User_service.GetUser(username)
		if err != nil {
			log.Printf(err.Error())
			return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
		}
		return c.JSON(common.NewSuccessResponseGetData(res, username, 1))
	} else {
		log.Printf("username empty")
		return c.JSON(common.NewBadRequestResponseWithMessage(errors.New("Username empty").Error()))
	}
}

func (controller *UsersController) GetAllUsersController(c echo.Context) error {
	res, err := controller.User_service.GetAllUser()
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseGetData(res, "All Users Data", len(res)))
}

func (controller *UsersController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Up",
	})
}
