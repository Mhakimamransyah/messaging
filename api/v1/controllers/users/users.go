package controllers

import (
	"messaging/api/common"
	"messaging/api/v1/controllers/users/response"
	"messaging/business"
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
		return c.JSON(common.NewBadRequestResponse())
	}
	return c.JSON(common.NewSuccessCreated())
}

func (controller *UsersController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {

		return c.JSON(common.NewBadRequestResponseWithMessage(business.ErrInvalidRequest.Error()))
	}
	user, err := controller.User_service.Login(username, password)
	if err != nil {

		return c.JSON(common.NewBadRequestResponse())
	}
	response := response.LoginResponse{ID: user.ID, Name: user.Name, Email: user.Email, RoleId: user.RolesId,
		Phone: user.Phone, Username: user.Username, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *UsersController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Up",
	})
}
