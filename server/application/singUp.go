package application

import (
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/db"
	"github.com/watariRyo/go-echo-redis/server/model/request"
	"github.com/watariRyo/go-echo-redis/server/model/response"

	"github.com/labstack/echo/v4"
)

// TODO contextを渡さないようにする
func SignUp(c echo.Context, signUpRequest *request.SignUpRequest) error {
	u := db.GetUser(&db.User{
		Name: signUpRequest.Name,
	})
	// Name重複はエラー
	if u.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "Name already exists",
		}
	}

	user := new(db.User)
	user.Name = signUpRequest.Name
	user.Password = signUpRequest.Password
	db.CreateUser(user)

	responseJSON := response.SignUpResponse{
		Message: "SignUp Success",
	}

	return c.JSON(http.StatusOK, responseJSON)
}
