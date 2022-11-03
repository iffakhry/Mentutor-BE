package delivery

import (
	"be12/mentutor/features/login"
	"be12/mentutor/utils/helper"
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
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		cnv := ToDomain(input)
		res, token, err := h.authUsecase.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err))
		}
		res.Token = token
		return c.JSON(http.StatusOK, helper.SuccessResponse("login successful", ToResponse(res, "login")))
	}
}
