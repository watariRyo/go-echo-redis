package application

import (
	"log"
	"net/http"
	"time"

	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/db"
	"github.com/watariRyo/go-echo-redis/server/helper"
	"github.com/watariRyo/go-echo-redis/server/model"
	"github.com/watariRyo/go-echo-redis/server/model/request"
	"github.com/watariRyo/go-echo-redis/server/model/response"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var signingKey = []byte(conf.SIGNING_KEY)

func Login(c echo.Context, loginRequest *request.LoginRequest) error {
	user := db.GetUser(&db.User{
		Name: loginRequest.Name,
	})

	if loginRequest.Name != user.Name || loginRequest.Password != user.Password { // FormとDBのデータを比較
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Name or Password",
		}
	}

	// セッション変数に値を付与
	session := helper.GetSession(c)
	session.Values["username"] = user.Name
	session.Values["auth"] = true
	if err := sessions.Save(c.Request(), c.Response()); err != nil {
		log.Fatal("Failed save session", err)
	}

	// JWT（Json Web Token）の処理
	claims := &model.JWTCustomClaims{
		UID:  user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}
	session.Values["token"] = signedToken

	responseJSON := response.LoginResponse{
		Name:  user.Name,
		Token: signedToken,
	}

	return c.JSON(http.StatusOK, responseJSON)
}

var JWTConfig = middleware.JWTConfig{
	Claims:     &model.JWTCustomClaims{},
	SigningKey: signingKey,
}
