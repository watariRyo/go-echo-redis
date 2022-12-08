package handler

import (
	"log"
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/domain"
	"github.com/watariRyo/go-echo-redis/server/model/request"

	"github.com/labstack/echo/v4"
)

type SignUpHandler struct {
	signUpDomain domain.SignUpDomain
}

func NewSignUpHandler(signUpDomain domain.SignUpDomain) SignUpHandler {
	return SignUpHandler{signUpDomain}
}

func (signUpHandler SignUpHandler) SignUp(c echo.Context) error {
	signUpRequest := new(request.SignUpRequest)

	if err := c.Bind(signUpRequest); err != nil {
		log.Printf("err %v", err.Error())
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Can not bind sign-up form",
		}
	}

	// 一旦nullだけ
	if signUpRequest.Name == "" || signUpRequest.Password == "" {
		log.Printf("err %v", "Invalid Name or Password")
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Invalid Name or Password",
		}
	}

	return signUpHandler.signUpDomain.SignUp(c, signUpRequest)
}
