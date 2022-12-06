package handler

import (
	"github.com/watariRyo/go-echo-redis/server/domain"

	"github.com/labstack/echo/v4"
)

func TestHandler(c echo.Context) error {
	return domain.Test(c)
}
