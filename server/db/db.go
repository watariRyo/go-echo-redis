package db

import (
	"log"

	"github.com/watariRyo/go-echo-redis/server/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	dbEndpoint = conf.DB_USER + ":" + conf.DB_PASS + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	var err error

	dsn := dbEndpoint
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	// Migration
	db.AutoMigrate(&User{})
}
