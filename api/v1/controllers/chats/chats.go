package chats

import (
	"messaging/api/common"
	"messaging/business/chats"
	"strconv"

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
	err, res := controller.Chats_service.SendMessage(specs)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.ChatCreated(res))
}

func (controller *ChatsController) ListChatController(c echo.Context) error {
	id_user, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Invalid parameter"))
	}
	res, err := controller.Chats_service.ListChat(id_user)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	query := make(map[string]interface{})
	query["id_user"] = id_user
	return c.JSON(common.NewSuccessResponseGetData(res, query, len(res)))
}

func (controller *ChatsController) ReadChatController(c echo.Context) error {
	id_user, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Invalid parameter"))
	}
	id_group, err := strconv.Atoi(c.Param("id_group"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Invalid parameter"))
	}
	res, err := controller.Chats_service.ReadChat(id_user, id_group)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	query := make(map[string]interface{})
	query["id_user"] = id_user
	query["id_group"] = id_group
	return c.JSON(common.NewSuccessResponseGetData(res, query, len(res)))
}
