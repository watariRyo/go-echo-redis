package handler

import (
	"github.com/watariRyo/go-echo-redis/server/domain"

	"github.com/labstack/echo/v4"
)

type TestHandler struct {
	testDomain domain.TestDomain
}

func NewTestHandler() TestHandler {
	return TestHandler{
		testDomain: domain.NewTestDomain(),
	}
}

func (testHandler TestHandler) Test(c echo.Context) error {
	return testHandler.testDomain.Test(c)
}
