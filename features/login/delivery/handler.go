package delivery

import (
	"be12/mentutor/features/login"
	"be12/mentutor/utils/helper"
	"errors"
	"log"
	"net/http"
	"strings"

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
			log.Print(err)
			if strings.Contains(err.Error(), "email not found") == true {
				errx := errors.New("user not found")
				return c.JSON(http.StatusNotFound, helper.FailedResponse(errx.Error()))
			} else {
				errx := errors.New("Invalid Input From Client")
				return c.JSON(http.StatusBadRequest, helper.FailedResponse(errx.Error()))
			}
		}

		if res.Role == "mentee"{
			// err := helper.Mailer()
			// if err != nil {
			// 	log.Print(err.Error(), " ERROR SEND MAIL")
			// }
		}

		res.Token = token
		return c.JSON(http.StatusOK, helper.SuccessResponse("login successful", ToResponse(res, "login")))
	}
}
