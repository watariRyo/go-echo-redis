package handler

import (
	"log"
	"net/http"

	"go-echo-redis/application"
	"go-echo-redis/model/request"

	"github.com/labstack/echo/v4"
)

func SignUpHandler(c echo.Context) error {
	signUpRequest := new(request.SignUpRequest)

	if err := c.Bind(signUpRequest); err != nil {
		log.Printf("err %v", err.Error())
		return &echo.HTTPError{
			Code: http.StatusInternalServerError,
			Message: "Can not bind sign-up form",
		}
	}

	// 一旦nullだけ
	if signUpRequest.Name == "" || signUpRequest.Password == "" {
		log.Printf("err %v", "Invalid Name or Password")
		return &echo.HTTPError{
			Code: http.StatusInternalServerError,
			Message: "Invalid Name or Password",
		}
	}

	return application.SignUp(c, signUpRequest)
}