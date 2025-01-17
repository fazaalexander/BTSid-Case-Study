package routes

import (
	"github.com/fazaalexander/BTSid-Case-Study.git/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	return e
}
