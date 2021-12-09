package chats

import (
	"log"
	"messaging/business"
	"messaging/business/users"
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
	User_repo  users.Repository
}

func InitChatService(repository Repository, repo_user users.Repository) *ChatService {
	return &ChatService{
		Chats_repo: repository,
		User_repo:  repo_user,
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

func (service *ChatService) ListChat(id_users int) ([]*ChatsList, error) {
	res, err := service.Chats_repo.GetChatsGroup(id_users)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, business.ErrInvalidRequest
	}
	var chatList []*ChatsList
	for key, _ := range res {
		value := **&res[key]
		user := users.Users{}
		if value.From_id_users == id_users {
			// gunakan to_id sebagai profile lawan chat
			tmp_user, err := service.User_repo.GetUserById(value.To_id_users)
			if err != nil {
				return nil, business.ErrInternalServerError
			}
			user = *tmp_user
		} else {
			// gunakan from_id sebagai profile lawan chat
			tmp_user, err := service.User_repo.GetUserById(value.From_id_users)
			if err != nil {
				return nil, business.ErrInternalServerError
			}
			user = *tmp_user
		}
		_chatList := ChatsList{
			IDGroup:        value.IDGroup,
			Chat_list_name: user.Name,
			Last_messages:  value.Messages,
			Unread:         service.Chats_repo.CountUnread(id_users, value.IDGroup),
		}
		chatList = append(chatList, &_chatList)
	}
	return chatList, nil
}

func (service *ChatService) ReadChat(id_users, id_groups int) ([]*ReadList, error) {
	res, err := service.Chats_repo.GetChatDetail(id_users, id_groups)
	if err != nil {
		log.Printf("%s", err)
		return nil, business.ErrInvalidRequest
	}

	var list_chat_read []*ReadList
	for _, data := range res {
		read := ReadList{
			ID:           data.ID,
			Type_chat:    GetType(id_users, data),
			Replies_chat: RepliesChat(data, res, id_users),
			Messages:     data.Messages,
			IsRead:       data.IsRead,
		}
		list_chat_read = append(list_chat_read, &read)
	}

	err = service.Chats_repo.UpdateRead(id_users, id_groups)
	if err != nil {
		log.Printf("%s", err)
	}
	return list_chat_read, nil
}
