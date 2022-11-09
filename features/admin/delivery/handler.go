package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/admin"
	"be12/mentutor/middlewares"
	"be12/mentutor/utils/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	e.POST("/admin/classes", handler.AddNewClass(), middleware.JWT([]byte(config.SECRET_JWT)))           // ADD NEW CLASS
	e.GET("/admin/classes", handler.GetAllClass(), middleware.JWT([]byte(config.SECRET_JWT)))            // GET ALL CLASS
	e.PUT("/admin/users/:id_user", handler.UpdateUserAdmin(), middleware.JWT([]byte(config.SECRET_JWT))) // UPDATE DATA USER BY ADMIN
	e.GET("/admin/users/:id_user", handler.GetSingleUser(), middleware.JWT([]byte(config.SECRET_JWT)))   // GET SINGLE PROFILE OTHER USER
	e.PUT("/admin/classes/:id_class", handler.UpdateClass(), middleware.JWT([]byte(config.SECRET_JWT)))	// UPDATE CLASS
	e.DELETE("/admin/classes/:id_class", handler.DeleteClass(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.DELETE("/admin/users/:id_user", handler.DeleteUser(), middleware.JWT([]byte(config.SECRET_JWT)))
}

func (ad *AdminDelivery) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		var input RegisterFormat

		if err := c.Bind(&input); err != nil {
			log.Print("error bind")
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		cnv := ToDomain(input)

		res, err := ad.adminUsecase.AddUser(cnv, role)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		return c.JSON(http.StatusCreated, helper.SuccessResponse("Register Success", ToResponse(res)))
	}
}

func (ad *AdminDelivery) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)


		resMentee, resMentor, err := ad.adminUsecase.GetAllUser(role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		return c.JSON(http.StatusOK, helper.SuccessResponse("success get all users", ToResponseUserArray(resMentee, resMentor)))
	}
}

func (ad *AdminDelivery) AddNewClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		var input AddClassFormat

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		cnvInput := ToDomainClass(input)
		res, err := ad.adminUsecase.AddNewClass(cnvInput, role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		return c.JSON(http.StatusCreated, helper.SuccessResponse("Success created", ToResponseAddClass(res)))
	}
}

func (ad *AdminDelivery) GetAllClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		log.Print("INI DI HANDLER")

		res, err := ad.adminUsecase.GetAllClass(role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Success Get All Class", ToResponseClassArray(res)))
	}
}

func (ad *AdminDelivery) UpdateUserAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		var input UpdateUserFormat

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		file, fileheader, err := c.Request().FormFile("images")
		if err == nil {
			res, err := helper.UploadFotoProfile(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			input.Images = res
		} 

		cnv := ToDomainUpdateUser(input)
		id := c.Param("id_user")
		cnvId, _ := strconv.Atoi(id)
		cnv.IdUser = uint(cnvId)

		res, err := ad.adminUsecase.UpdateUserAdmin(cnv, role)
		if  err != nil && strings.Contains(err.Error(), "user") == true{
			log.Print(err)
			return c.JSON(http.StatusNotFound, helper.FailedResponse("User Not Found"))
		}  else if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, helper.SuccessResponse("update profile successful", ToResponseUpdateUser(res)))
	}
}

func (ad *AdminDelivery) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)


		id := c.Param("id_user")
		cnv, _ := strconv.Atoi(id)

		err := ad.adminUsecase.DeleteUser(uint(cnv), role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Delete Success", ToResponseDeleteUser(cnv)))
	}
}

func (ad *AdminDelivery) GetSingleUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)


		id := c.Param("id_user")
		cnv, _ := strconv.Atoi(id)

		res, err := ad.adminUsecase.GetSingleUser(uint(cnv), role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Success Get Profile", ToResponseGetUser(res)))
	}
}

func (ad *AdminDelivery) UpdateClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		var input UpdateClassFormat

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		id := c.Param("id_class")
		cnvId, _ := strconv.Atoi(id)

		log.Print(input)
		cnvData := ToDomainUpdateClass(input)
		cnvData.IdClass = uint(cnvId)
		log.Print(cnvData)
		res, err := ad.adminUsecase.UpdateClass(cnvData, role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, helper.SuccessResponse("Update Class Successful", ToResponseUpdateClass(res)))
	}
}

func (ad *AdminDelivery) DeleteClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		id := c.Param("id_class")
		cnv, _ := strconv.Atoi(id)

		err := ad.adminUsecase.DeleteClass(uint(cnv), role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
			return c.JSON(http.StatusOK, helper.SuccessResponse("Success Delete Class", map[string]int{"id_class": cnv}))
	}
}