package application

import (
	"go-echo-redis/db"
	"go-echo-redis/model/request"
	"go-echo-redis/model/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context, signUpRequest *request.SignUpRequest) error {
	u := db.GetUser(&db.User{
		Name: signUpRequest.Name,
	})
	// Name重複はエラー
	if u.ID != 0 {
		return &echo.HTTPError{
			Code: http.StatusConflict,
			Message: "Name already exists",
		}
	}

	user := new(db.User)
	user.Name = signUpRequest.Name
	user.Password = signUpRequest.Password
	db.CreateUser(user)

	responseJSON := response.SignUpResponse {
		Message: "SignUp Success",
	}

	return c.JSON(http.StatusOK, responseJSON)
}