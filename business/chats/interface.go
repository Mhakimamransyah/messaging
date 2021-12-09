package chats

type Service interface {
	SendMessage(chat_specs ChatsSpec) (error, interface{})
	ListChat(id_users int) ([]*ChatsList, error)
	ReadChat(id_users, id_groups int) ([]*ReadList, error)
}

type Repository interface {
	CreateChats(chats *Chats) (error, interface{})
	GetChatDetail(id_users, id_group int) ([]*Chats, error)
	GetChatsGroup(id_user int) ([]*Chats, error)
	CountUnread(id_users, id_group int) int
	UpdateRead(id_users, id_group int) error
}
