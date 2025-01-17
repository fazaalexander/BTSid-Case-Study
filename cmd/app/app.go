package app

import (
	"github.com/fazaalexander/BTSid-Case-Study.git/cmd/routes"
	"github.com/fazaalexander/BTSid-Case-Study.git/common"
	"github.com/fazaalexander/BTSid-Case-Study.git/database/mysql"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	handler := common.Handler{}

	router := routes.StartRoute(handler)

	return router
}
