package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/admin"
	"go/constant"

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
		var input 
	}
}