//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/domain"
	"github.com/watariRyo/go-echo-redis/server/handler"
	"github.com/watariRyo/go-echo-redis/server/repository"
	"gorm.io/gorm"
)

func InitializeService(configApiServer *conf.ApiServer, db *gorm.DB) *Server {
	wire.Build(
		repository.NewUserRepository,
		domain.NewLoginDomain,
		handler.NewLoginHandler,
		domain.NewSignUpDomain,
		handler.NewSignUpHandler,
		domain.NewTestDomain,
		handler.NewTestHandler,
		NewServer,
	)
	return &Server{}
}
