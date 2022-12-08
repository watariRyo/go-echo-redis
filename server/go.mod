module github.com/watariRyo/go-echo-redis/server

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/wire v0.5.0
	github.com/gorilla/sessions v1.2.1
	github.com/labstack/echo-contrib v0.13.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/rbcervilla/redisstore/v8 v8.1.0
	github.com/spf13/viper v1.14.0
	gorm.io/driver/mysql v1.4.1
	gorm.io/gorm v1.24.0
)
