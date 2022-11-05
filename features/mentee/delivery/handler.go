package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/mentee"
	"be12/mentutor/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MenteeDelivery struct {
	MenteeUsecase mentee.UseCaseInterface
}

func New(e *echo.Echo, usecase mentee.UseCaseInterface) {

	handler := MenteeDelivery{
		MenteeUsecase: usecase,
	}

	e.PUT("/update/:id_user", handler.UpdateProfile()) // UPDATE PROFILE USER
	e.POST("/forum", handler.AddStatus(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.GET("/forum", handler.SelectAll(), middleware.JWT([]byte(config.SECRET_JWT)))
}

func (md *MenteeDelivery) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat

		IdUser, _, _ := middlewares.ExtractToken(c)
		cnvInput := ToEntity(input)

		res, err := md.MenteeUsecase.UpdateProfile(uint(IdUser), cnvInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("Something Error In Server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success update profile", res))
	}
}

func (md *MenteeDelivery) AddStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Request
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("Something Error In Server"))
		}
		id, _, role := middlewares.ExtractToken(c)

		if role != "mentee" {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}

		cnv := ToDomain(input)

		res, errposts := md.MenteeUsecase.InsertStatus(cnv, id)
		if errposts != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success add status", ToResponse(res)))
	}
}

func (md *MenteeDelivery) SelectAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _, _ := middlewares.ExtractToken(c)
		if id < 1 {
			return c.JSON(http.StatusNotFound, FailedResponse("Invalid Input From Client"))
		}
		res, err := md.MenteeUsecase.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all status", ToCoreArray(res)))

	}
}
