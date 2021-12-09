package chats

import (
	"time"
)

type Chats struct {
	ID              int
	IDGroup         int
	From_id_users   int
	To_id_users     int
	Replies_id_chat int
	IsRead          int
	Messages        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type ChatsList struct {
	IDGroup        int    `json:"id_group"`
	Chat_list_name string `json:"name"`
	Phone          string `json:"phone"`
	Username       string `json:"username"`
	Last_messages  string `json:"messages"`
	Unread         int    `json:"unread_messages"`
}

type ReadList struct {
	ID           int       `json:"id_chat"`
	Type_chat    string    `json:"type_message"`
	Replies_chat *ReadList `json:"replies_from_chat"`
	IsRead       int       `json:"is_read"`
	Messages     string    `json:"messages"`
	Date         time.Time `json:"send_at"`
}

func GetType(id_user int, chat *Chats) string {
	if id_user == chat.From_id_users {
		return "Sender"
	} else {
		return "Receiver"
	}
}

func RepliesChat(chat *Chats, list_chat []*Chats, id_user int) *ReadList {
	replies := ReadList{}
	if chat.Replies_id_chat != 0 {
		for _, data := range list_chat {
			if data.ID == chat.Replies_id_chat {
				replies.ID = data.ID
				replies.Date = data.CreatedAt
				replies.Type_chat = GetType(id_user, data)
				replies.IsRead = data.IsRead
				replies.Messages = data.Messages
				break
			}
		}
		return &replies
	} else {
		return nil
	}
}

type ChatDetailList struct {
}

func NewChats(chats_specs *ChatsSpec) *Chats {
	return &Chats{
		IDGroup:         chats_specs.IdGroup,
		From_id_users:   chats_specs.From_id_users,
		To_id_users:     chats_specs.To_id_users,
		Replies_id_chat: chats_specs.Replies_id_chat,
		Messages:        chats_specs.Messages,
		IsRead:          0,
		CreatedAt:       time.Now(),
	}
}
