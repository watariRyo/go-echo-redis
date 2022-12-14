package main

import (
	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/repository"
	"github.com/watariRyo/go-echo-redis/server/server"
)

func main() {
	// TODO
	// server.goをmainに変換
	// server設定部分を分離
	// wire導入

	// Injection
	//login := handler.NewLoginHandler(domain.NewLoginDomain(repository.NewUserRepository()))
	//signUp := handler.NewSignUpHandler(domain.NewSignUpDomain(repository.NewUserRepository()))
	//test := handler.NewTestHandler(domain.NewTestDomain())

	// echoインスタンスを生成
	//e := echo.New()

	cfg, err := conf.Load()
	if err != nil {
		panic(err)
	}
	s := server.InitializeService(&cfg.ApiServer, repository.LoadClient())

	println(cfg.Db.Host)
	println(cfg.ApiServer.CorsOrigins)

	// サーバ起動
	s.Run()
}
