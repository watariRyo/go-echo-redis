package repository

import (
	"github.com/watariRyo/go-echo-redis/server/db"
	"github.com/watariRyo/go-echo-redis/server/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{
		db: db.LoadClient(),
	}
}

func (ur UserRepository) GetUser(u *model.User) model.User {
	var user model.User
	ur.db.Where(u).First(&user)
	return user
}

func (ur UserRepository) CreateUser(u *model.User) {
	ur.db.Create(u)
}
