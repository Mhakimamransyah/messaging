package chats

import (
	"messaging/api/common"
	"messaging/business/chats"

	"github.com/labstack/echo/v4"
)

type ChatsController struct {
	Chats_service chats.Service
}

func InitChatsController(service chats.Service) *ChatsController {
	return &ChatsController{
		Chats_service: service,
	}
}

func (controller *ChatsController) SendMessagesController(c echo.Context) error {
	specs := chats.ChatsSpec{}
	c.Bind(&specs)
	err := controller.Chats_service.SendMessage(specs)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessCreated())
}

func (controller *ChatsController) ListChatController(c echo.Context) error {
	return nil
}

func (controller *ChatsController) ReadChatController(c echo.Context) error {
	return nil
}
