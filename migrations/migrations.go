package migrations

import (
	"messaging/repositories/chats"
	"messaging/repositories/users"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.UsersTable{},
		&chats.ChatsTable{},
	)
}
