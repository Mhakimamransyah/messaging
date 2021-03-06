package api

import (
	apikey "messaging/api/middleware/apikey"
	ControllersChats "messaging/api/v1/controllers/chats"
	ControllersUsers "messaging/api/v1/controllers/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, user_controller *ControllersUsers.UsersController, chat_controller *ControllersChats.ChatsController) {

	users := e.Group("v1/users")
	users.Use(apikey.ApiKey())
	users.POST("", user_controller.Register)
	users.GET("", user_controller.GetAllUsersController)
	users.GET("/:username", user_controller.GetUsersByUsername)

	chats := e.Group("v1/chats")
	chats.Use(apikey.ApiKey())
	chats.POST("", chat_controller.SendMessagesController)
	chats.GET("/:id_user", chat_controller.ListChatController)
	chats.GET("/:id_user/read/:id_group", chat_controller.ReadChatController)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "Welcome to Messaging Service",
			"description": "Rakamin Academy Backend Mini Projects",
			"developer": map[string]interface{}{
				"name":             "M.Hakim Amransyah",
				"linkedin":         "https://www.linkedin.com/in/hakim-amr/",
				"personal-website": "http://mhakimamransyah.site/",
			},
			"repository": "https://github.com/Mhakimamransyah/messaging",
		})
	})
}
