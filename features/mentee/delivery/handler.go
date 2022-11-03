package delivery

import (
	"be12/mentutor/features/mentee"
	"be12/mentutor/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	MenteeUsecase mentee.UseCaseInterface
}

func New(e *echo.Echo, usecase mentee.UseCaseInterface) {

	handler := MenteeDelivery{
		MenteeUsecase: usecase,
	}

	e.PUT("/update/:id_user", handler.UpdateProfile()) // UPDATE PROFILE USER
}

func (md *MenteeDelivery) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat

		id, _ := middlewares.ExtractToken(c)
		cnvInput := ToEntity(input)
		
		res, err := md.MenteeUsecase.UpdateProfile(uint(id), cnvInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("Something Error In Server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success update profile", res))
	}
}