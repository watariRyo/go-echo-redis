package handler

import (
	"log"
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/helper"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func LogoutHandler(c echo.Context) error {
	session := helper.GetSession(c)
	// ログアウト
	session.Values["auth"] = false
	// セッション削除
	session.Options.MaxAge = -1
	if err := sessions.Save(c.Request(), c.Response()); err != nil {
		log.Fatal("Failed cannot delete session", err)
	}
	return c.JSON(http.StatusOK, "logout")
}
