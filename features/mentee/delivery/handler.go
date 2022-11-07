package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/mentee"
	"be12/mentutor/middlewares"
	"be12/mentutor/utils/helper"
	"errors"
	"log"
	"net/http"
	"strconv"

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
	e.POST("/forum/:id", handler.AddComment(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.POST("/mentees/submission/:id", handler.AddSub(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.POST("/mentees/sub/:id", handler.AddSubMis(), middleware.JWT([]byte(config.SECRET_JWT)))

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
		file, fileheader, err := c.Request().FormFile("images")
		if err == nil {
			res, err := helper.UploadStatusImages(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		}

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
		res, resC, err := md.MenteeUsecase.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}

		// return c.JSON(http.StatusOK, SuccessResponse("success get all status", ToCoreArray(res)))
		return c.JSON(http.StatusOK, SuccessResponse("success get all status", ToCoreArray(res, resC)))

	}
}
func (md *MenteeDelivery) AddComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var comment CommentFormat
		id_status := c.Param("id")

		if err := c.Bind(&comment); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("Invalid Input From Client"))
		}

		idUser, _, role := middlewares.ExtractToken(c)
		idCnv, _ := strconv.Atoi(id_status)
		idStatus := uint(idCnv)
		comment.IdStatus = idStatus
		comment.ID_User = uint(idUser)
		data := ToDomainComments(comment)
		log.Print(data)
		if role == "admin" {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}
		res, err1 := md.MenteeUsecase.Insert(data)
		if err1 != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("error from server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success insert comment", ToResponseComments(res)))

	}
}

func (md *MenteeDelivery) AddSub() echo.HandlerFunc {
	return func(c echo.Context) error {
		var submission SubFormat
		idtasks := c.Param("id")

		if err := c.Bind(&submission); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("Invalid Input From Client"))
		}

		if err := c.Bind(&submission); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("Invalid Input From Client"))
		}
		file, fileheader, err := c.Request().FormFile("file")
		if err == nil {
			res, err := helper.UploadFileSubmisiion(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			submission.File = res
		}

		idUser, _, role := middlewares.ExtractToken(c)
		idCnv, _ := strconv.Atoi(idtasks)
		IdTask := uint(idCnv)
		submission.ID_Tasks = IdTask
		submission.ID_Mentee = uint(idUser)
		data := ToDomainSub(submission)
		log.Print(data)
		if role != "mentee" {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}
		res, err1 := md.MenteeUsecase.InsertSub(data)
		if err1 != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("error from server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success insert submission", ToResponseSub(res)))

	}
}
func (md *MenteeDelivery) AddSubMis() echo.HandlerFunc {
	return func(c echo.Context) error {
		var submission SubFormat
		idtasks := c.Param("id")

		if err := c.Bind(&submission); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("Invalid Input From Client"))
		}
		file, fileheader, err := c.Request().FormFile("file")
		if err == nil {
			res, err := helper.UploadFileSubmisiion(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			submission.File = res
		}
		idUser, _, role := middlewares.ExtractToken(c)
		idCnv, _ := strconv.Atoi(idtasks)
		IdTask := uint(idCnv)
		submission.ID_Tasks = IdTask
		submission.ID_Mentee = uint(idUser)
		data := ToDomainSub(submission)
		log.Print(data)
		if role != "mentee" {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}
		res, err1 := md.MenteeUsecase.InsertSubmis(int(IdTask), data)
		if err1 != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("error from server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success insert submission", ToResponseSub(res)))

	}
}
