package model

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
    UID  int    `json:"uid"`
    Name string `json:"name"`
    jwt.StandardClaims
}