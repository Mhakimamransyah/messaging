package chats

import (
	"log"
	"messaging/business"
	"messaging/util/validator"
)

type ChatsSpec struct {
	IdGroup         int    `form:"id_group" json:"id_group"`
	From_id_users   int    `form:"id_users_sender" json:"id_users_sender" validate:"required"`
	To_id_users     int    `form:"id_users_receiver" json:"id_users_receiver" validate:"required"`
	Replies_id_chat int    `form:"id_chats_replies" json:"id_chats_replies"`
	Messages        string `form:"messages" json:"messages" validate:"required"`
}

type ChatService struct {
	Chats_repo Repository
}

func InitChatService(repository Repository) *ChatService {
	return &ChatService{
		Chats_repo: repository,
	}
}

func (service *ChatService) SendMessage(chat_specs ChatsSpec) error {
	err := validator.GetValidator().Struct(chat_specs)
	if err != nil {
		return business.ErrInvalidSpec
	}
	chats := NewChats(&chat_specs)
	err = service.Chats_repo.CreateChats(chats)
	if err != nil {
		log.Printf("%s", err.Error())
		return business.ErrInvalidRequest
	}
	return nil
}

func (service *ChatService) ListChat(id_users int) ([]*Chats, error) {
	return nil, nil
}

func (service *ChatService) ReadChat(id_users, id_groups int) ([]*Chats, error) {
	return nil, nil
}
