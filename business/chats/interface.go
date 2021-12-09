package chats

type Service interface {
	SendMessage(chat_specs ChatsSpec) error
	ListChat(id_users int) ([]*ChatsList, error)
	ReadChat(id_users, id_groups int) ([]*ReadList, error)
}

type Repository interface {
	CreateChats(chats *Chats) error
	GetChatDetail(id_users, id_group int) ([]*Chats, error)
	GetChatsGroup(id_user int) ([]*Chats, error)
	CountUnread(id_users, id_group int) int
	UpdateRead(id_users, id_group int) error
}
