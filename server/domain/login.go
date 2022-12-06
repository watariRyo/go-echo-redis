package domain

import (
	"github.com/watariRyo/go-echo-redis/server/repository"
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/helper"
	"github.com/watariRyo/go-echo-redis/server/model"
	"github.com/watariRyo/go-echo-redis/server/model/request"
	"github.com/watariRyo/go-echo-redis/server/model/response"

	"github.com/labstack/echo/v4"
)

type LoginDomain struct {
	userRepository repository.UserRepository
}

func NewLoginDomain() LoginDomain {
	return LoginDomain{
		userRepository: repository.NewUserRepository(),
	}
}

func (loginDomain LoginDomain) Login(c echo.Context, loginRequest *request.LoginRequest) error {
	user := loginDomain.userRepository.GetUser(&model.User{
		Name: loginRequest.Name,
	})

	if loginRequest.Name != user.Name || loginRequest.Password != user.Password { // FormとDBのデータを比較
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Name or Password",
		}
	}

	signedToken := helper.JwtCreateToken(user)

	helper.SessionGrantValueToAuth(c, user.Name, signedToken)

	responseJSON := response.LoginResponse{
		Name:  user.Name,
		Token: signedToken,
	}

	return c.JSON(http.StatusOK, responseJSON)
}

// var JWTConfig = middleware.JWTConfig{
// 	Claims:     &model.JWTCustomClaims{},
// 	SigningKey: log,
// }
