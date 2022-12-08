package repository

import (
	"github.com/watariRyo/go-echo-redis/server/model"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ur UserRepository) GetUser(u *model.User) model.User {
	var user model.User
	db.Where(u).First(&user)
	return user
}

func (ur UserRepository) CreateUser(u *model.User) {
	db.Create(u)
}
