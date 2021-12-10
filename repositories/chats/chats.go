package chats

import (
	"fmt"
	"log"
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
	IDGroup         int              `gorm:"id_group;id;default:0"`
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

func ConvertChatsTableToChats(chat *ChatsTable) *chats.Chats {
	return &chats.Chats{
		ID:              chat.ID,
		From_id_users:   chat.From_id_users,
		To_id_users:     chat.To_id_users,
		Replies_id_chat: chat.Replies_id_chat,
		IDGroup:         chat.IDGroup,
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

func (repos *ChatRepository) CreateChats(chats *chats.Chats) (error, interface{}) {

	chatsTable := ConvertChatsToChatsTables(chats)
	var id_group interface{}

	if chatsTable.IDGroup > 0 {
		// already chatting before

		// if replies, check if replies id exist
		replied_to_chat := ChatsTable{}
		err := repos.DB.Where("id = ?", chats.ID).First(&replied_to_chat).Error
		if err != nil {
			return err, nil
		}
		err = repos.DB.Save(chatsTable).Error
		if err != nil {
			return err, nil
		}
		id_group = chatsTable.IDGroup
	} else {
		// first time chatting

		query := fmt.Sprintf("INSERT INTO chats( id_group, from_id_users, to_id_users,messages,replies_id_chat, created_at) SELECT COALESCE(MAX(id_group), 0) + 1,%d,%d,'%s',%d, NOW() FROM chats", chatsTable.From_id_users, chatsTable.To_id_users,
			chats.Messages, chatsTable.Replies_id_chat)
		res := repos.DB.Exec(query).Model(&chatsTable)
		if res.Error != nil {
			return res.Error, nil
		}
		id_group = chatsTable.IDGroup
	}
	res := make(map[string]interface{})
	res["id_group"] = id_group
	return nil, res
}

func (repos *ChatRepository) UpdateRead(id_users, id_group int) error {
	chat_tables := []ChatsTable{}
	err := repos.DB.Model(&chat_tables).Where("id_group = ? AND to_id_users = ?", id_group, id_users).Update("is_read", 1).Error
	if err != nil {
		return err
	}
	return nil
}

func (repos *ChatRepository) GetChatDetail(id_users, id_group int) ([]*chats.Chats, error) {
	chats_table_group := []ChatsTable{}
	err := repos.DB.Where("id_group = ?", id_group).Order("chats.created_at asc").Find(&chats_table_group).Error
	if err != nil {
		log.Printf("%s", err)
	}

	var list_chat []*chats.Chats
	for i, _ := range chats_table_group {
		chat := ConvertChatsTableToChats(&chats_table_group[i])
		list_chat = append(list_chat, chat)
	}
	return list_chat, nil
}

func (repos *ChatRepository) GetChatsGroup(id_user int) ([]*chats.Chats, error) {
	// GET , GROUP AND ORDER ID GROUP
	chats_table_group := []ChatsTable{}
	err := repos.DB.Where("from_id_users = ? OR to_id_users = ?", id_user, id_user).Order("created_at desc").Find(&chats_table_group).Error
	if err != nil {
		return nil, err
	}

	// GROUP HERE
	unique_group := []ChatsTable{}
	for key, data := range chats_table_group {
		if key == 0 {
			unique_group = append(unique_group, data)
		} else {
			exists := false
			for _, data2 := range unique_group {
				if data2.IDGroup == data.IDGroup {
					exists = true
					break
				}
			}
			if exists == false {
				unique_group = append(unique_group, data)
			}
		}
	}

	// GET AND ORDER LAST CHAT MESSAGE
	chats_group := []ChatsTable{}
	for _, group := range unique_group {
		var chat_group_ ChatsTable
		repos.DB.Where("id_group = ?", group.IDGroup).Order("chats.created_at desc").First(&chat_group_)
		chats_group = append(chats_group, chat_group_)
	}

	// CONVERT AND SEND
	var list_chat []*chats.Chats
	for i, _ := range chats_group {
		chat := ConvertChatsTableToChats(&chats_group[i])
		list_chat = append(list_chat, chat)
	}
	return list_chat, nil
}

func (repos *ChatRepository) CountUnread(id_users, id_group int) int {
	chats_table_group := []ChatsTable{}
	repos.DB.Where("id_group = ? AND to_id_users = ? AND is_read = ?", id_group, id_users, 0).Find(&chats_table_group)
	return len(chats_table_group)
}
