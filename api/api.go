package api

import "github.com/labstack/echo/v4"

type Api interface {
	AddFile() echo.HandlerFunc
}
