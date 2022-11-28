package handler

import (
	"log"
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/application"
	"github.com/watariRyo/go-echo-redis/server/model/request"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	loginRequest := new(request.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		log.Printf("err %v", err.Error())
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Can not bind login form",
		}
	}

	return application.Login(c, loginRequest)
}
