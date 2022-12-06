package domain

import (
	"github.com/watariRyo/go-echo-redis/server/repository"
	"net/http"

	"github.com/watariRyo/go-echo-redis/server/model"
	"github.com/watariRyo/go-echo-redis/server/model/request"
	"github.com/watariRyo/go-echo-redis/server/model/response"

	"github.com/labstack/echo/v4"
)

type SignUpDomain struct {
	userRepository repository.UserRepository
}

func NewSignUpDomain() SignUpDomain {
	return SignUpDomain{
		userRepository: repository.NewUserRepository(),
	}
}

func (signUpDomain SignUpDomain) SignUp(c echo.Context, signUpRequest *request.SignUpRequest) error {
	u := signUpDomain.userRepository.GetUser(&model.User{
		Name: signUpRequest.Name,
	})
	// Name重複はエラー
	if u.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "Name already exists",
		}
	}

	user := new(model.User)
	user.Name = signUpRequest.Name
	user.Password = signUpRequest.Password
	signUpDomain.userRepository.CreateUser(user)

	responseJSON := response.SignUpResponse{
		Message: "SignUp Success",
	}

	return c.JSON(http.StatusOK, responseJSON)
}
