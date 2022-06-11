package api

import "github.com/labstack/echo/v4"

type Api interface {
	// file
	AddFile() echo.HandlerFunc
	GetLastFile() echo.HandlerFunc

	// region
	GetRegionPredictList() echo.HandlerFunc
}