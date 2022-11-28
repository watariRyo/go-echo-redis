package main

import (
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/handler"
	"github.com/watariRyo/go-echo-redis/server/helper"
	"github.com/watariRyo/go-echo-redis/server/model"
	"github.com/watariRyo/go-echo-redis/server/model/response"

	"github.com/golang-jwt/jwt"
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
	e.POST("/signUp", handler.SignUpHandler)
	e.POST("/login", handler.LoginHandler)
	e.POST("/logout", handler.LogoutHandler)

	r := e.Group("/api")
	// JWT検証
	config := middleware.JWTConfig{
		Claims:     &model.JWTCustomClaims{},
		SigningKey: []byte(conf.SIGNING_KEY),
	}
	r.Use(middleware.JWTWithConfig(config))
	// ルート（認証必要（/api/**））
	r.GET("/", hello)

	// サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}

// 認証サンプル
func hello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JWTCustomClaims)

	s := helper.GetSession(c)

	responseJSON := response.HelloResponse{
		UID:  claims.UID,
		Name: s.Values["username"].(string),
	}
	return c.JSON(http.StatusOK, responseJSON)
}
