package chats

type Service interface {
	SendMessage(chat_specs ChatsSpec) error
	ListChat(id_users int) ([]*Chats, error)
	ReadChat(id_users, id_groups int) ([]*Chats, error)
}

type Repository interface {
	CreateChats(chats *Chats) error
	GetListChats(id_users int) ([]*Chats, error)
	GetChatDetail(id_users, id_group int) ([]*Chats, error)
}
