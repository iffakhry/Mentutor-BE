package delivery

import (
	"be12/mentutor/features/login"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	authUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := AuthDelivery{
		authUsecase: usecase,
	}

	e.POST("/login", handler.Login())

}

func (h *AuthDelivery) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "failed Bind Data",
			})
		}

		cnv := ToDomain(input)
		res, token, err := h.authUsecase.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed Login",
			})
		}
		res.Token = token
		return c.JSON(http.StatusOK, SuccessResponseWithData("login successful", ToResponse(res, "login")))
	}
}
