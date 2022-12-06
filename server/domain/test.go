package domain

import (
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/helper"
	"github.com/watariRyo/go-echo-redis/server/model"
	"github.com/watariRyo/go-echo-redis/server/model/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JWTCustomClaims)

	s := helper.GetSession(c)

	responseJSON := response.HelloResponse{
		UID:  claims.UID,
		Name: s.Values["username"].(string),
	}
	return c.JSON(http.StatusOK, responseJSON)
}
