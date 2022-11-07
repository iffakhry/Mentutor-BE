package delivery

import (
	"be12/mentutor/features/mentor"
	"log"
	"time"
)

type UpdateUserFormat struct {
	ID uint
	IdClass uint
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Images string `json:"images" form:"images"`
}

type TaskRequest struct {
	Title string `form:"title"`
	Description string `form:"description"`
	Images string `form:"images"`
	File string `form:"file"` 
	DueDate string `form:"due_date"`
	IdMentor uint
	IdClass uint
}

func ToDomainUpdateUser(data UpdateUserFormat) mentor.UserCore {
	return mentor.UserCore{
		IdUser: data.ID,
		IdClass: data.IdClass,
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		Images: data.Images,
	}
}

func ToDomainTask(data TaskRequest) mentor.TaskCore {
	layoutFormat := "2006-01-02 15:04:05.999 MST"
	value := data.DueDate + " " + "23:59:59.000 WIB"
	dateRes, _ := time.Parse(layoutFormat, value)
	log.Print(dateRes)
	return mentor.TaskCore{
		Description: data.Description,
		IdMentor: data.IdMentor,
		IdClass: data.IdClass,
		File: data.File,
		DueDate: dateRes,
		Title: data.Title,
	}
}

