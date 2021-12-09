package chats_test

import (
	"errors"
	"messaging/business/chats"
	chat_mocks "messaging/business/chats/mocks"
	"messaging/business/users"
	users_mocks "messaging/business/users/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	ID_Users        = 5
	ID              = 3
	IDGroup         = 1
	From_id_users   = 1
	To_id_users     = 2
	Replies_id_chat = 1
	IsRead          = 0
	Messages        = "hallo there"
)

var (
	// Data
	chatData     chats.Chats
	chatSpec     chats.ChatsSpec
	chatDataList []*chats.Chats
	usersData    users.Users
	// Method
	chat_repo    chat_mocks.Repository
	user_repo    users_mocks.Repository
	chat_service chats.Service
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestSendMessage(t *testing.T) {

	t.Run("Expects Invalid Chat Specs", func(t *testing.T) {
		failed_specs := chats.ChatsSpec{
			IdGroup:       1,
			From_id_users: 1,
		}
		err, res := chat_service.SendMessage(failed_specs)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("Expects Failed Insert Chat in Database", func(t *testing.T) {
		chat_repo.On("CreateChats", mock.AnythingOfType("*chats.Chats")).Return(errors.New(""), nil).Once()
		err, res := chat_service.SendMessage(chatSpec)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("Expects Success send new chats", func(t *testing.T) {
		chat_repo.On("CreateChats", mock.AnythingOfType("*chats.Chats")).Return(nil, chatData).Once()
		err, res := chat_service.SendMessage(chatSpec)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
}

func TestListChat(t *testing.T) {
	t.Run("Expects error while get chat group", func(t *testing.T) {
		chat_repo.On("GetChatsGroup", mock.AnythingOfType("int")).Return(nil, errors.New("Error")).Once()
		res, err := chat_service.ListChat(ID_Users)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("Expects error read profile partner", func(t *testing.T) {
		chat_repo.On("GetChatsGroup", ID_Users).Return(chatDataList, nil).Once()
		user_repo.On("GetUserById", mock.AnythingOfType("int")).Return(nil, errors.New("Error")).Once()
		res, err := chat_service.ListChat(ID_Users)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	// t.Run("Expects error read profile partner", func(t *testing.T) {
	// 	chat_repo.On("GetChatsGroup", From_id_users).Return(chatDataList, nil).Once()
	// 	user_repo.On("GetUserById", mock.AnythingOfType("int")).Return(nil, errors.New("Error")).Once()
	// 	res, err := chat_service.ListChat(ID_Users)
	// 	assert.NotNil(t, err)
	// 	assert.Nil(t, res)
	// })

	t.Run("Expects success", func(t *testing.T) {
		chat_repo.On("GetChatsGroup", ID_Users).Return(chatDataList, nil).Once()
		user_repo.On("GetUserById", mock.AnythingOfType("int")).Return(&usersData, nil).Once()
		chat_repo.On("CountUnread", ID_Users, IDGroup).Return(2).Once()
		res, err := chat_service.ListChat(ID_Users)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
}

func TestReadChat(t *testing.T) {
	t.Run("Expects failed get chat detail", func(t *testing.T) {
		chat_repo.On("GetChatDetail", ID_Users, IDGroup).Return(nil, errors.New("Error")).Once()
		res, err := chat_service.ReadChat(ID_Users, IDGroup)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
	t.Run("Expects failed update read", func(t *testing.T) {
		chat_repo.On("GetChatDetail", ID_Users, IDGroup).Return(chatDataList, nil).Once()
		chat_repo.On("UpdateRead", ID_Users, IDGroup).Return(errors.New("Error")).Once()
		res, err := chat_service.ReadChat(ID_Users, IDGroup)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func setup() {

	chat_service = chats.InitChatService(&chat_repo, &user_repo)

	chatData = chats.Chats{
		ID:              ID,
		IDGroup:         IDGroup,
		From_id_users:   From_id_users,
		To_id_users:     To_id_users,
		Replies_id_chat: Replies_id_chat,
		IsRead:          IsRead,
		Messages:        Messages,
	}

	chatSpec = chats.ChatsSpec{
		IdGroup:         IDGroup,
		From_id_users:   From_id_users,
		To_id_users:     To_id_users,
		Replies_id_chat: Replies_id_chat,
		Messages:        Messages,
	}

	usersData = users.Users{
		ID:       1,
		Name:     "M.Hakim Amransyah",
		Username: "mhakim",
		Phone:    "081271286874",
	}

	chatDataList = append(chatDataList, &chatData)

}
