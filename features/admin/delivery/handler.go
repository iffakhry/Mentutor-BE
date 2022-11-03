package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/admin"
	"be12/mentutor/utils/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AdminDelivery struct {
	adminUsecase admin.UsecaseInterface
}

func New(e *echo.Echo, usecase admin.UsecaseInterface) {

	handler := AdminDelivery{
		adminUsecase: usecase,
	}

	e.POST("/admin/users", handler.AddUser(), middleware.JWT([]byte(config.SECRET_JWT)))                 // ADD NEW USER
	e.GET("/admin/users", handler.GetAllUser(), middleware.JWT([]byte(config.SECRET_JWT)))               // GET ALL USER
	e.POST("/admin/classes", handler.AddNewClass(), middleware.JWT([]byte(config.SECRET_JWT)))           //ADD NEW CLASS
	e.GET("/admin/classes", handler.GetAllClass(), middleware.JWT([]byte(config.SECRET_JWT)))            //GET ALL CLASS
	e.PUT("/admin/users/:id_user", handler.UpdateUserAdmin(), middleware.JWT([]byte(config.SECRET_JWT))) // UPDATE DATA USER BY ADMIN
	e.DELETE("/admin/mentee/:id_user", handler.DeleteUserMentee(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.DELETE("/admin/mentor/:id_user", handler.DeleteUserMentor(), middleware.JWT([]byte(config.SECRET_JWT)))
}

func (ad *AdminDelivery) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat

		if err := c.Bind(&input); err != nil {
			log.Print("error bind")
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

func (ad *AdminDelivery) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		resMentee, resMentor, err := ad.adminUsecase.GetAllUser(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		return c.JSON(http.StatusOK, helper.SuccessResponse("success get all users", ToResponseUserArray(resMentee, resMentor)))
	}
}

func (ad *AdminDelivery) AddNewClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddClassFormat

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		cnvInput := ToDomainClass(input)
		err := ad.adminUsecase.AddNewClass(cnvInput, c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		return c.JSON(http.StatusOK, helper.SuccessResponseNoData("Success created"))
	}
}

func (ad *AdminDelivery) GetAllClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ad.adminUsecase.GetAllClass(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Success Get All Class", ToResponseClassArray(res)))
	}
}

func (ad *AdminDelivery) UpdateUserAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateUserFormat

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		cnv := ToDomainUpdateUser(input)
		id := c.Param("id_user")
		cnvId, _ := strconv.Atoi(id)
		cnv.IdUser = uint(cnvId)
		res, err := ad.adminUsecase.UpdateUserAdmin(cnv, c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("update profile successful", ToResponseUpdateUser(res)))
	}
}

func (ad *AdminDelivery) DeleteUserMentee() echo.HandlerFunc {
	return func(c echo.Context) error {

		id:= c.Param("id_user")
		cnv, _ := strconv.Atoi(id)

		err := ad.adminUsecase.DeleteUserMentee(uint(cnv), c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponseNoData("Delete Success"))
	}
}

func (ad *AdminDelivery) DeleteUserMentor() echo.HandlerFunc {
	return func(c echo.Context) error {

		id:= c.Param("id_user")
		cnv, _ := strconv.Atoi(id)

		err := ad.adminUsecase.DeleteUserMentor(uint(cnv), c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponseNoData("Delete Success"))
	}
}

