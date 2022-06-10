package server

import (
	"hack/api"
	"hack/internal/app/config"
	"hack/internal/app/store"

	_ "hack/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type server struct {
	router *echo.Echo
	store  store.Store
	config config.Config
	api    api.Api
}

// NewServer инициализируем сервер
func NewServer(store store.Store, config config.Config, api api.Api) *server {
	s := &server{
		router: echo.New(),
		store:  store,
		config: config,
		api:    api,
	}

	// Конфигурируем роутинг
	s.configureRouter()
	return s
}

// Start Включаем прослушивание HTTP протокола
func (s *server) Start(address string) error {
	return s.router.Start(address)
}

// configureRouter Объявляем список доступных роутов
func (s *server) configureRouter() {
	s.router.Use(
		middleware.CORS(),
		middleware.RequestID(),
		middleware.Logger(),
	)
	s.router.GET("/readyz", s.handleReadyz)
	s.router.GET("/statusz", s.handleHealthz)
	s.router.GET("/swagger/*", echoSwagger.WrapHandler)
	s.router.GET("/ws", s.handleWS)
	s.router.GET("/test", s.hello)
}
