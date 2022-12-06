package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/model"
)

var secret = ""

func init() {
	cfg, err := conf.Load()
	if err != nil {
		panic(err)
	}
	secret = cfg.Jwt.Key
}

func JwtCreateToken(user model.User) string {
	// JWT（Json Web Token）の処理
	claims := &model.JWTCustomClaims{
		UID:  user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return signedToken
}
