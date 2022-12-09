package helper

import (
	"gorm.io/gorm"
)

func Transaction(tx *gorm.DB, f func() error) error {
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
