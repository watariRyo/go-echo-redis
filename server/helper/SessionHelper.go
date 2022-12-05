package helper

import (
	"context"
	"log"

	"github.com/watariRyo/go-echo-redis/server/conf"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v8"
)

var redisEndpoint = ""
var sessionKey = ""

func init() {
	cfg, err := conf.Load()
	if err != nil {
		panic(err)
	}
	redisEndpoint = cfg.Redis.Host + ":" + cfg.Redis.Port
	sessionKey = cfg.Redis.Key
}

func GetSession(c echo.Context) *sessions.Session {
	client := redis.NewClient(&redis.Options{
		Addr:     redisEndpoint,
		Password: "",
		DB:       0,
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal("Failed cannot connect redis", err)
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		MaxAge:   600,
		HttpOnly: true,
	})
	session, err := store.Get(c.Request(), sessionKey)
	if err != nil {
		log.Fatal("Failed cannot get session", err)
	}
	return session
}
