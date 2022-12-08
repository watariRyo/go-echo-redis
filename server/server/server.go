package server

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/watariRyo/go-echo-redis/server/conf"
	"github.com/watariRyo/go-echo-redis/server/domain"
	"github.com/watariRyo/go-echo-redis/server/handler"
	"github.com/watariRyo/go-echo-redis/server/model"
	"github.com/watariRyo/go-echo-redis/server/repository"
	"net/http"
	"strings"
)

func NewServer(
	conf *conf.ApiServer,
	userRepository repository.UserRepository,
	loginDomain domain.LoginDomain,
	loginHandler handler.LoginHandler,
	signUpDomain domain.SignUpDomain,
	signUpHandler handler.SignUpHandler,
	testDomain domain.TestDomain,
	testHandler handler.TestHandler,
) *Server {
	server := &Server{
		echo.New(),
		conf,
		userRepository,
		loginDomain,
		loginHandler,
		signUpDomain,
		signUpHandler,
		testDomain,
		testHandler,
	}
	server.setRouting()
	server.setCors()
	return server
}

type Server struct {
	e               *echo.Echo
	configApiServer *conf.ApiServer
	userRepository  repository.UserRepository
	loginDomain     domain.LoginDomain
	loginHandler    handler.LoginHandler
	signUpDomain    domain.SignUpDomain
	signUpHandler   handler.SignUpHandler
	testDomain      domain.TestDomain
	testHandler     handler.TestHandler
}

func (server *Server) setRouting() {
	// middleware設定
	server.e.Use(middleware.Logger())
	server.e.Use(middleware.Recover())
	server.e.Use(middleware.Gzip())

	cfg, err := conf.Load()
	if err != nil {
		panic(err)
	}
	// セッション
	server.e.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.Jwt.Key))))

	// ルート（認証不要）
	server.e.POST("/echo/signUp", server.signUpHandler.SignUp)
	server.e.POST("/echo/login", server.loginHandler.Login)
	server.e.POST("/echo/logout", handler.LogoutHandler)

	r := server.e.Group("/echo/api")
	// JWT検証
	config := middleware.JWTConfig{
		Claims:     &model.JWTCustomClaims{},
		SigningKey: []byte(cfg.Jwt.Key),
	}
	r.Use(middleware.JWTWithConfig(config))
	// ルート（認証必要（/api/**））
	r.GET("/", server.testHandler.Test)
}

func (server *Server) setCors() {
	var corsOrigins []string
	if len(server.configApiServer.CorsOrigins) != 0 {
		corsOrigins = strings.Split(server.configApiServer.CorsOrigins, ",")
	}
	server.e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     corsOrigins,
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}),
	)
	// corsで403を返却させる
	server.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Originヘッダの中身を取得
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			// 許可しているOriginの中でリクエストヘッダのOriginと一致すれば処理継続
			for _, o := range corsOrigins {
				if origin == o {
					return next(c)
				}
			}
			// 一致しない場合は403
			return &echo.HTTPError{
				Code:    http.StatusForbidden,
				Message: "cors forbidden",
			}
		}
	})
}

func (server *Server) Run() {
	server.e.Logger.Fatal(server.e.Start(":" + server.configApiServer.Port))
}
