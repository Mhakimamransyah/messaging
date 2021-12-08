package users

import (
	"messaging/business/users"
	"time"

	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func InitRepository(DB *gorm.DB) *UsersRepository {
	return &UsersRepository{
		DB: DB,
	}
}

type UsersTable struct {
	gorm.Model
	ID        int       `gorm:"id;primaryKey:autoIncrement"`
	Name      string    `gorm:"name;not null;type:varchar(100);uniqueIndex:Name"`
	Email     string    `gorm:"email;not null;type:varchar(100);uniqueIndex:Email"`
	Username  string    `gorm:"username;not null;type:varchar(100);uniqueIndex:Username"`
	Password  string    `gorm:"password;not null;type:text"`
	Phone     string    `gorm:"phone;not null;type:varchar(100);uniqueIndex:Phone"`
	RolesId   int       `gorm:"roles_id;not null;type:bigint;"`
	CreatedAt time.Time `gorm:"created_at;type:datetime;default:null"`
	UpdatedAt time.Time `gorm:"updated_at;type:datetime;default:null"`
	DeletedAt time.Time `gorm:"deleted_at;type:datetime;default:null"`
}

func ConvertUsersToUsersTable(users *users.Users) *UsersTable {
	return &UsersTable{
		ID:        users.ID,
		Name:      users.Name,
		Email:     users.Email,
		Username:  users.Username,
		Password:  users.Password,
		Phone:     users.Phone,
		RolesId:   users.RolesId,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
		DeletedAt: users.DeletedAt,
	}
}

func ConvertUserTablesToUsers(user_table *UsersTable) *users.Users {
	return &users.Users{
		ID:        user_table.ID,
		Name:      user_table.Name,
		Email:     user_table.Email,
		Phone:     user_table.Phone,
		Username:  user_table.Username,
		Password:  user_table.Password,
		RolesId:   user_table.RolesId,
		CreatedAt: user_table.CreatedAt,
		UpdatedAt: user_table.UpdatedAt,
		DeletedAt: user_table.DeletedAt,
	}
}

func (repo *UsersRepository) CreateUser(user *users.Users) error {
	userTable := ConvertUsersToUsersTable(user)
	err := repo.DB.Save(userTable).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UsersRepository) Login(username, password string) (*users.Users, error) {
	user_table := UsersTable{}
	err := repo.DB.Where("username = ?", username).First(&user_table).Error
	if err != nil {
		return nil, err
	}
	return ConvertUserTablesToUsers(&user_table), nil
}

func (repo *UsersRepository) Get(username string) (*users.Users, error) {
	return nil, nil
}
func (repo *UsersRepository) Update(user *users.Users) error {
	return nil
}
func (repo *UsersRepository) Delete(user *users.Users) error {
	return nil
}
