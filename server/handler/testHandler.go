package handler

import (
	"github.com/watariRyo/go-echo-redis/server/domain"

	"github.com/labstack/echo/v4"
)

type TestHandler struct {
	testDomain domain.TestDomain
}

func NewTestHandler(testDomain domain.TestDomain) TestHandler {
	return TestHandler{testDomain}
}

func (testHandler TestHandler) Test(c echo.Context) error {
	return testHandler.testDomain.Test(c)
}
