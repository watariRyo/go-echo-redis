package main

import (
	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/repository"
	"github.com/watariRyo/go-echo-redis/server/server"
)

func main() {

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
