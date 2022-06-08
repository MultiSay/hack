package server

import (
	"hack/internal/app/model"
	"hack/internal/app/system"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handleHealthz Возвращаем статус приложения
func (s *server) handleHealthz(ctx echo.Context) error {
	err := system.Healthz()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, model.Status{
		Status: model.StatusOK,
	})
}

// handleReadyz Возвращаем статус приложения
func (s *server) handleReadyz(ctx echo.Context) error {
	err := system.Readyz()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, model.Status{
		Status: model.StatusOK,
	})
}

func (s *server) ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			log.Printf("%s", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}
