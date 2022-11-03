package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	adminDelivery "be12/mentutor/features/admin/delivery"
	adminData "be12/mentutor/features/admin/repository"
	adminUsecase "be12/mentutor/features/admin/services"

	loginDelivery "be12/mentutor/features/login/delivery"
	loginData "be12/mentutor/features/login/repository"
	loginUsecase "be12/mentutor/features/login/services"

	menteeDelivery "be12/mentutor/features/mentee/delivery"
	menteeData "be12/mentutor/features/mentee/repository"
	menteeUsecase "be12/mentutor/features/mentee/services"

	statusDelivery "be12/mentutor/features/forum/delivery"
	statusData "be12/mentutor/features/forum/repository"
	statusUsecase "be12/mentutor/features/forum/service"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	adminDataFactory := adminData.New(db)
	adminUsecaseFactory := adminUsecase.New(adminDataFactory)
	adminDelivery.New(e, adminUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	menteeDataFactory := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeDataFactory)
	menteeDelivery.New(e, menteeUsecaseFactory)

	statusData := statusData.New(db)
	statusUsecaseFactory := statusUsecase.New(statusData)
	statusDelivery.New(e, statusUsecaseFactory)

}
