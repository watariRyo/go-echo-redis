package repository

import (
	"github.com/watariRyo/go-echo-redis/server/model"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ur UserRepository) GetUser(u *model.User, tx *gorm.DB) model.User {
	var user model.User
	if tx != nil {
		tx.Where(u).First(&user)
	} else {
		db.Where(u).First(&user)
	}
	return user
}

func (ur UserRepository) CreateUser(u *model.User, tx *gorm.DB) {
	tx.Create(u)
}
