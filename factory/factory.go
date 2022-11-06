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

	mentorDelivery "be12/mentutor/features/mentor/delivery"
	mentorData "be12/mentutor/features/mentor/repository"
	mentorUsecase "be12/mentutor/features/mentor/services"

	menteeDelivery "be12/mentutor/features/mentee/delivery"
	menteeData "be12/mentutor/features/mentee/repository"
	menteeUsecase "be12/mentutor/features/mentee/services"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	adminDataFactory := adminData.New(db)
	adminUsecaseFactory := adminUsecase.New(adminDataFactory)
	adminDelivery.New(e, adminUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	mentorDataFactory := mentorData.New(db)
	mentorUsecaseFactory := mentorUsecase.New(mentorDataFactory)
	mentorDelivery.New(e, mentorUsecaseFactory)

	menteeDataFactory := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeDataFactory)
	menteeDelivery.New(e, menteeUsecaseFactory)

}
