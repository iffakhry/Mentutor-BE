package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/admin"
	"be12/mentutor/utils/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AdminDelivery struct {
	adminUsecase admin.UsecaseInterface
}

func New(e *echo.Echo, usecase admin.UsecaseInterface) {

	handler :=AdminDelivery{
		adminUsecase: usecase,
	}

	e.POST("/admin/users", handler.AddUser(), middleware.JWT([]byte(config.SECRET_JWT)))
}

func (ad *AdminDelivery) AddUser() echo.HandlerFunc{
	return func(c echo.Context) error {
		var input RegisterFormat
		
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		cnv := ToDomain(input)

		res, err := ad.adminUsecase.AddUser(cnv, c)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		return c.JSON(http.StatusCreated, helper.SuccessResponse("Register Success", ToResponse(res)))	
	}
}