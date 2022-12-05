package main

import (
	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/handler"
	"github.com/watariRyo/go-echo-redis/server/model"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// echoインスタンスを生成
	e := echo.New()

	// middleware設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// セッション
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(conf.SIGNING_KEY))))

	// ルート（認証不要）
	e.POST("/echo/signUp", handler.SignUpHandler)
	e.POST("/echo/login", handler.LoginHandler)
	e.POST("/echo/logout", handler.LogoutHandler)

	r := e.Group("/echo//api")
	// JWT検証
	config := middleware.JWTConfig{
		Claims:     &model.JWTCustomClaims{},
		SigningKey: []byte(conf.SIGNING_KEY),
	}
	r.Use(middleware.JWTWithConfig(config))
	// ルート（認証必要（/api/**））
	r.GET("/", handler.TestHandler)

	// サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}
