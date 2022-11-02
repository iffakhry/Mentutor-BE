package delivery

import (
	"be12/mentutor/features/mentee"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	MenteeUsecase mentee.UseCaseInterface
}

func New(e *echo.Echo, usecase mentee.UseCaseInterface) {

	handler := MenteeDelivery{
		MenteeUsecase: usecase,
	}

	e.POST("/users/update", handler.UpdateProfile())
}

func (md *MenteeDelivery) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		var id uint

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}

		idCnv, err := strconv.Atoi(c.Param("id_user"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}
		id = uint(idCnv)

		cnvInput := ToEntity(input)

		res, err := md.MenteeUsecase.UpdateProfile(id, cnvInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("Something Error In Server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success update profile", res))
	}
}