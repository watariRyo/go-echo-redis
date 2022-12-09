package repository

import (
	"log"

	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	cfg, err := conf.Load()
	if err != nil {
		panic(err)
	}

	dsn := cfg.Db.User + ":" + cfg.Db.Password + "@tcp(" + cfg.Db.Host + ":" + cfg.Db.Port + ")/" + cfg.Db.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println(err)
	}

	// Migration
	db.AutoMigrate(&model.User{})
}

func LoadClient() *gorm.DB {
	return db
}
