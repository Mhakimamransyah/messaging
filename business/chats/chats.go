package chats

import "time"

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
