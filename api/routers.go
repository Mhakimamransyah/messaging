package api

import (
	ControllersChats "messaging/api/v1/controllers/chats"
	ControllersUsers "messaging/api/v1/controllers/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, user_controller *ControllersUsers.UsersController, chat_controller *ControllersChats.ChatsController) {

	users := e.Group("v1/users")
	users.POST("", user_controller.Register)
	users.GET("/:username", user_controller.GetUsersByUsername)

	chats := e.Group("v1/chats")
	chats.POST("", chat_controller.SendMessagesController)
	chats.GET("/:id_users", chat_controller.ListChatController)
	chats.GET("/:id_users/read/:id_group", chat_controller.ReadChatController)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})
}
