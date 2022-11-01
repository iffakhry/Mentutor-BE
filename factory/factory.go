package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userDelivery "be12/mentutor/features/user/delivery"
	userData "be12/mentutor/features/user/repository"
	userUsecase "be12/mentutor/features/user/services"

	loginDelivery "be12/mentutor/features/login/delivery"
	loginData "be12/mentutor/features/login/repository"
	loginUsecase "be12/mentutor/features/login/services"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)
}
