package model

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
	DriverID   int
	DriverRole int
	ShopID     int
}
