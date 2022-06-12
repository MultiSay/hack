package server

import (
	"hack/api"
	"hack/internal/app/config"
	"hack/internal/app/store"
	"hack/internal/app/websocket"
	"net/http"

	_ "hack/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type server struct {
	router *echo.Echo
	store  store.Store
	config config.Config
	v1     api.Api
	ws     *websocket.WS
}

// NewServer инициализируем сервер
func NewServer(store store.Store, config config.Config, api api.Api, ws *websocket.WS) *server {
	s := &server{
		router: echo.New(),
		store:  store,
		config: config,
		v1:     api,
		ws:     ws,
	}

	// Конфигурируем роутинг
	s.configureRouter()
	return s
}

// Start Включаем прослушивание HTTP протокола
func (s *server) Start(address string) error {
	return s.router.Start(address)
}

func setResponseACAOHeaderFromRequest(req http.Request, resp echo.Response) {
	resp.Header().Set(echo.HeaderAccessControlAllowOrigin,
		req.Header.Get(echo.HeaderAllow))
}

func ACAOHeaderOverwriteMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Before(func() {
			setResponseACAOHeaderFromRequest(*ctx.Request(), *ctx.Response())
		})
		return next(ctx)
	}
}

// configureRouter Объявляем список доступных роутов
func (s *server) configureRouter() {
	s.router.Use(
		ACAOHeaderOverwriteMiddleware,
		middleware.RequestID(),
		middleware.Logger(),
		middleware.CORS(),
	)
	s.router.GET("/readyz", s.handleReadyz)
	s.router.GET("/statusz", s.handleHealthz)
	s.router.GET("/swagger/*", echoSwagger.WrapHandler)
	s.router.GET("/ws", s.handleWS)
	s.router.GET("/test", s.hello)
	v1 := s.router.Group("/v1")
	{
		v1.POST("/file", s.v1.AddFile())
		v1.GET("/file", s.v1.GetLastFile())
		v1.GET("/lead", s.v1.GetLeadList())
		v1.GET("/region/predict", s.v1.GetRegionPredictList())
	}
}
