package helper

import (
	"github.com/watariRyo/go-echo-redis/server/repository"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = repository.LoadClient()
}

func Transaction(f func() error) error {
	// repository側でdbを使ってしまっているのでトランザクションがかからない
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := f(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
