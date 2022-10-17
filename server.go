package main

import (
	"go-echo-redis/conf"
	"go-echo-redis/handler"
	"go-echo-redis/model"
	"go-echo-redis/model/response"
	"net/http"

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
    responseJSON := response.HelloResponse{
        UID:  claims.UID,
        Name: claims.Name,
    }
    return c.JSON(http.StatusOK, responseJSON)
}