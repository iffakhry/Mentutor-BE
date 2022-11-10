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
	"strings"

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

	e.POST("/forum", handler.AddStatus(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.GET("/forum", handler.SelectAll(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.POST("/forum/:id", handler.AddComment(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.POST("/mentees/submission/:id", handler.AddSub(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.GET("/mentees/tasks", handler.GetAllTasks(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.POST("/gmail", handler.GmailRequest())
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
		res, resC, resMntr, err := md.MenteeUsecase.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}

		// return c.JSON(http.StatusOK, SuccessResponse("success get all status", ToCoreArray(res)))
		return c.JSON(http.StatusOK, SuccessResponse("success get all status", ToCoreArray(res, resC, resMntr)))

	}
}

func (md *MenteeDelivery) GetAllTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, idClass, role := middlewares.ExtractToken(c)
		if id < 1 {
			return c.JSON(http.StatusNotFound, FailedResponse("Invalid Input From Client"))
		}

		res, err := md.MenteeUsecase.GetTask(uint(idClass), role)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}

		// return c.JSON(http.StatusOK, SuccessResponse("success get all status", ToCoreArray(res)))
		return c.JSON(http.StatusOK, SuccessResponse("success get all status", tasksResponse(res)))

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
			log.Print(err1)
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

		file, fileheader, err := c.Request().FormFile("file")
		if err == nil {
			res, err := helper.UploadFileSubmission(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			submission.File = res
		}

		idUser, _, role := middlewares.ExtractToken(c)
		idCnv, _ := strconv.Atoi(idtasks)
		submission.ID_Tasks = uint(idCnv)
		submission.ID_Mentee = uint(idUser)
		data := ToDomainSub(submission)

		if role != "mentee" {
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}
		res, err1 := md.MenteeUsecase.InsertSub(data)
		if err1 != nil {
			log.Print(err1)
			if strings.Contains(err1.Error(), "due date") {
				return c.JSON(http.StatusBadRequest, FailedResponse("Submit melewati due date"))
			}
			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
		}

		// GET TOKEN UNTUK AUTH GMAIL
		token , err := helper.GetToken()
		if err != nil {
			log.Print(err.Error())
			return c.JSON(http.StatusInternalServerError, FailedResponse("Failed get token gmail"))
		}
		res.Token = token
		return c.JSON(http.StatusCreated, SuccessResponse("success insert submission", ToResponseSub(res)))

	}
}

func (md *MenteeDelivery) GmailRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data GmailFormat
		
		c.Bind(&data)

		log.Print(data)
		helper.SendGmail()
		return nil
	}
}
