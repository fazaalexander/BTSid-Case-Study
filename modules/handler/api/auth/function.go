package auth

import (
	"net/http"

	ue "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
	"github.com/labstack/echo/v4"
)

func (ah *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.RegisterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := ah.authUseCase.Register(request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "Register Success",
			"Status":  http.StatusOK,
		})
	}
}

func (ah *AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.LoginRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		data, err := ah.authUseCase.Login(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "Login Success",
			"Data":    data,
			"Status":  http.StatusOK,
		})
	}
}
