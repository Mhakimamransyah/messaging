package chats

import (
	"fmt"
	"messaging/business/chats"
	"messaging/repositories/users"
	"time"

	"gorm.io/gorm"
)

type ChatRepository struct {
	DB *gorm.DB
}

func InitRepository(DB *gorm.DB) *ChatRepository {
	return &ChatRepository{
		DB: DB,
	}
}

type ChatsTable struct {
	gorm.Model
	ID              int              `gorm:"id;primaryKey:autoIncrement;unique"`
	IDGroup         int              `gorm:"id_group;id;"`
	From_id_users   int              `gorm:"from_id_users;not null;"`
	To_id_users     int              `gorm:"to_id_users;"`
	Replies_id_chat int              `gorm:"replies_id_chat"`
	IsRead          int              `gorm:"isread;default:0"`
	Messages        string           `gorm:"messages;not null;type:text"`
	CreatedAt       time.Time        `gorm:"created_at;type:datetime;default:null"`
	UpdatedAt       time.Time        `gorm:"updated_at;type:datetime;default:null"`
	DeletedAt       time.Time        `gorm:"deleted_at;type:datetime;default:null"`
	Users_sender    users.UsersTable `gorm:"foreignKey:from_id_users"`
	Users_receiver  users.UsersTable `gorm:"foreignKey:to_id_users"`
}

func ConvertChatsToChatsTables(chat *chats.Chats) *ChatsTable {
	return &ChatsTable{
		IDGroup:         chat.IDGroup,
		From_id_users:   chat.From_id_users,
		To_id_users:     chat.To_id_users,
		Replies_id_chat: chat.Replies_id_chat,
		IsRead:          chat.IsRead,
		Messages:        chat.Messages,
		CreatedAt:       chat.CreatedAt,
		UpdatedAt:       chat.UpdatedAt,
		DeletedAt:       chat.DeletedAt,
	}
}

type Tabler interface {
	TableName() string
}

func (ChatsTable) TableName() string {
	return "chats"
}

func (repos *ChatRepository) CreateChats(chats *chats.Chats) error {
	chatsTable := ConvertChatsToChatsTables(chats)
	if chatsTable.IDGroup > 0 {
		// already chatting before
		err := repos.DB.Save(chatsTable).Error
		if err != nil {
			return err
		}
	} else {
		// first time chatting
		query := fmt.Sprintf("INSERT INTO chats( id_group, from_id_users, to_id_users,messages,replies_id_chat) SELECT MAX( id_group )+1,%d,%d,'%s',%d FROM chats", chatsTable.From_id_users, chatsTable.To_id_users,
			chats.Messages, chatsTable.Replies_id_chat)
		err := repos.DB.Exec(query).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (repos *ChatRepository) GetListChats(id_users int) ([]*chats.Chats, error) {
	return nil, nil
}

func (repos *ChatRepository) GetChatDetail(id_users, id_group int) ([]*chats.Chats, error) {
	return nil, nil
}
