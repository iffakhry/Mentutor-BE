package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/mentor"
	"be12/mentutor/middlewares"
	"be12/mentutor/utils/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MentorDelivery struct {
	mentorUsecase mentor.UsecaseInterface
}

func New(e *echo.Echo, usecase mentor.UsecaseInterface) {

	handler := MentorDelivery{
		mentorUsecase: usecase,
	}

	e.PUT("/users", handler.UpdateProfile(), middleware.JWT([]byte(config.SECRET_JWT)))               // UPDATE USER
	e.POST("/mentors/tasks", handler.AddTask(), middleware.JWT([]byte(config.SECRET_JWT)))            // UPDATE USER
	e.GET("/mentors/tasks", handler.GetAllTask(), middleware.JWT([]byte(config.SECRET_JWT)))          // GET ALL TASk
	e.GET("/mentors/tasks/:id_task", handler.GetTaskSub(), middleware.JWT([]byte(config.SECRET_JWT))) // GET TASK BY ID TASK
}

func (md *MentorDelivery) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateUserFormat

		IdUser, IdClass, role := middlewares.ExtractToken(c)
		input.IdClass = uint(IdClass)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		// CEK GAMBAR
		file, fileheader, err := c.Request().FormFile("images")
		if err == nil {
			res, err := helper.UploadFotoProfile(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		}

		input.ID = uint(IdUser)
		cnv := ToDomainUpdateUser(input)
		res, err := md.mentorUsecase.UpdateProfile(cnv, role)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, helper.SuccessResponse("success update profile", ToResponseUpdateUser(res)))
	}
}

func (md *MentorDelivery) AddTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		IdUser, IdClass, role := middlewares.ExtractToken(c)

		var input TaskRequest

		if err := c.Bind(&input); err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		// CEK GAMBAR
		file, fileheader, err := c.Request().FormFile("images")
		if err == nil {
			res, err := helper.UploadGambarTugas(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		}

		// CEK FILE
		file, fileheader, err = c.Request().FormFile("file")
		if err == nil {
			res, err := helper.UploadFileTugas(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		}

		input.IdClass = uint(IdClass)
		input.IdMentor = uint(IdUser)
		cnv := ToDomainTask(input)
		res, err := md.mentorUsecase.AddTask(cnv, role)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, helper.SuccessResponse("Success insert task", ToResponseAddTask(res)))
	}
}

func (md *MentorDelivery) GetAllTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		res, err := md.mentorUsecase.GetAllTask((role))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, helper.SuccessResponse("Success get all task", ToResponseGetAllTask(res)))
	}
}

func (md *MentorDelivery) GetTaskSub() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, role := middlewares.ExtractToken(c)

		idTask := c.Param("id_task")
		cnvId, _ := strconv.Atoi(idTask)

		resTask, resSub, err := md.mentorUsecase.GetTaskSub(uint(cnvId), role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, helper.SuccessResponse("success get single task", ToResponseSingleTask(resTask, resSub)))
	}
}