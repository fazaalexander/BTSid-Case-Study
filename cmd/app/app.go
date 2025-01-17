package app

import (
	"github.com/fazaalexander/BTSid-Case-Study.git/cmd/routes"
	"github.com/fazaalexander/BTSid-Case-Study.git/common"
	"github.com/fazaalexander/BTSid-Case-Study.git/database/mysql"

	authHandler "github.com/fazaalexander/BTSid-Case-Study.git/modules/handler/api/auth"
	authRepo "github.com/fazaalexander/BTSid-Case-Study.git/modules/repository/auth"
	authUseCase "github.com/fazaalexander/BTSid-Case-Study.git/modules/usecase/auth"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	authRepo := authRepo.New(mysql.DB)
	authUseCase := authUseCase.New(authRepo)
	authHandler := authHandler.New(authUseCase)

	handler := common.Handler{
		AuthHandler: authHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
