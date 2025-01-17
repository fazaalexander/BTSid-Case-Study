package auth

import "github.com/labstack/echo/v4"

func (ah *AuthHandler) RegisterRoutes(e *echo.Echo) {
	authGroup := e.Group("/v1/auth")
	authGroup.POST("/register", ah.Register())
	authGroup.POST("/login", ah.Login())
}
