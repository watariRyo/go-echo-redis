package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/domain"
	"github.com/watariRyo/go-echo-redis/server/handler"
	"github.com/watariRyo/go-echo-redis/server/model"
	"github.com/watariRyo/go-echo-redis/server/repository"
)

func main() {
	// TODO
	// server.goをmainに変換
	// server設定部分を分離
	// wire導入

	// Injection
	login := handler.NewLoginHandler(domain.NewLoginDomain(repository.NewUserRepository()))
	signUp := handler.NewSignUpHandler(domain.NewSignUpDomain(repository.NewUserRepository()))
	test := handler.NewTestHandler(domain.NewTestDomain())

	// echoインスタンスを生成
	e := echo.New()

	cfg, err := conf.Load()
	if err != nil {
		panic(err)
	}

	// middleware設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// セッション
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.Jwt.Key))))

	// ルート（認証不要）
	e.POST("/echo/signUp", signUp.SignUp)
	e.POST("/echo/login", login.Login)
	e.POST("/echo/logout", handler.LogoutHandler)

	r := e.Group("/echo/api")
	// JWT検証
	config := middleware.JWTConfig{
		Claims:     &model.JWTCustomClaims{},
		SigningKey: []byte(cfg.Jwt.Key),
	}
	r.Use(middleware.JWTWithConfig(config))
	// ルート（認証必要（/api/**））
	r.GET("/", test.Test)

	// サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}
