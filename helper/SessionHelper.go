package helper

import (
	"context"
	"go-echo-redis/conf"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v8"
)

const (
	redisEndpoint = conf.REDIS_HOST + ":" + conf.REDIS_PORT
	sessionKey = conf.SESSION_KEY
)

func GetSession(c echo.Context) *sessions.Session {
	client := redis.NewClient(&redis.Options{
		Addr: redisEndpoint,
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal("Failed cannot connect redis", err)
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		MaxAge: 600,
		HttpOnly: true,
	})
	session, err := store.Get(c.Request(), sessionKey)
	if err != nil {
		log.Fatal("Failed cannot get session", err)
	}
	return session
}
