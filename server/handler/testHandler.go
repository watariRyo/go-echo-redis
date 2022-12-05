package handler

import (
	"github.com/watariRyo/go-echo-redis/server/application"

	"github.com/labstack/echo/v4"
)

func TestHandler(c echo.Context) error {
	return application.Test(c)
}
