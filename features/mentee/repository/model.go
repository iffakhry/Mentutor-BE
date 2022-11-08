package repository

import (
	"be12/mentutor/features/mentee"
	"time"

	"gorm.io/gorm"
)

//MODEL MENTEE
type Mentee struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	Images   string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:enum('mentee');not null"`
	IdClass  uint
}

// MODEL STATUS
type Status struct {
	gorm.Model
	IdMentee uint   `json:"id_mentee" form:"Id_mentee"`
	Name     string `gorm:"->"`
	Images   string `json:"images" form:"images"`
	Caption  string `json:"caption" form:"caption"`
}

// MODEL KOMENTAR
type Comments struct {
	gorm.Model
	ID_User  uint   `json:"id_user" form:"id_user"`
	IdStatus uint   `json:"id_status" form:"id_status"`
	Caption  string `json:"caption" form:"caption"`
	Name     string `gorm:"<-:false"`
	Role     string `gorm:"<-:false"`
}

// MODEL SUBMISSION
type Submission struct {
	gorm.Model
	IdMentee uint
	IdTask   uint
	File     string
	Score    uint
	Title    string `gorm:"->"`
	Status  string
}

// MODEL TASK
type Task struct {
	gorm.Model
	IdClass     uint
	IdMentor    uint
	Images      string
	DueDate     *time.Time
	Score       uint   `gorm:"->"`
	Status      string `gorm:"->"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	File        string `gorm:"type:varchar(255);not null"`
}

// MENTEE
func FromEntity(data mentee.MenteeCore) Mentee {
	return Mentee{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Images:   data.Images,
	}
}

func ToEntity(id uint, data Mentee) mentee.MenteeCore {
	return mentee.MenteeCore{
		IdUser: id,
		Name:   data.Name,
		Email:  data.Email,
		Images: data.Images,
	}
}

func ToEntityMentee(data mentee.Status) Status {
	return Status{
		Model:    gorm.Model{ID: data.ID},
		IdMentee: data.IdMentee,
		Caption:  data.Caption,
		Images:   data.Images,
		// Comments: data.Comment,
	}
}

// STATUS
func toPostUser(dataPost Status) mentee.Status {

	dataPostCore := mentee.Status{
		ID:       dataPost.ID,
		IdMentee: dataPost.IdMentee,
		Images:   dataPost.Images,
		Caption:  dataPost.Caption,
		Name:     dataPost.Name,
		// Comment:  dataPost.Comments,
	}

	return dataPostCore

}

func ToCoreArray(status []Status) []mentee.Status {
	var res []mentee.Status
	for _, val := range status {
		res = append(res, mentee.Status{
			ID:        val.ID,
			Images:    val.Images,
			Caption:   val.Caption,
			IdMentee:  val.IdMentee,
			CreatedAt: val.CreatedAt,
		})
	}
	return res
}

// COMMENTS
func (comment *Comments) ToDomainComments(data Comments) mentee.CommentsCore {
	return mentee.CommentsCore{
		ID: data.ID,
		// ID_User:    data.ID_User,
		IdStatus:   data.IdStatus,
		Caption:    data.Caption,
		Role:       data.Role,
		Name:       data.Name,
		Created_At: data.CreatedAt,
	}
}

func ToEntityComent(data mentee.CommentsCore) Comments {
	return Comments{

		ID_User:  data.ID_User,
		IdStatus: data.IdStatus,
		Caption:  data.Caption,
		// Role:     data.Role,
		// Name:     data.Name,
	}
}

func FromEntityComment(data Comments) mentee.CommentsCore {
	return mentee.CommentsCore{
		ID:       data.ID,
		ID_User:  data.ID_User,
		IdStatus: data.IdStatus,
		Caption:  data.Caption,
		// Role:     data.Role,
		// Name:     data.Name,
	}
}

func toPostList(data []Status) []mentee.Status {
	var dataCore []mentee.Status
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, mentee.Status{ID: data[i].ID, Caption: data[i].Caption, Images: data[i].Images, Name: data[i].Name})
	}
	return dataCore
}
func toTaskList(data []Task) []mentee.Task {
	var dataCore []mentee.Task
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, mentee.Task{
			ID:          data[i].ID,
			Title:       data[i].Title,
			Images:      data[i].Images,
			Description: data[i].Description,
			File:        data[i].File,
			Score:       data[i].Score,
			DueDate:     *data[i].DueDate,
			Status:      data[i].Status,
		})
	}
	return dataCore
}
func ToComent(data []Comments) []mentee.CommentsCore {
	var dataCmn []mentee.CommentsCore
	for _, v := range data {
		dataCmn = append(dataCmn, mentee.CommentsCore{
			ID:       v.ID,
			ID_User:  v.ID_User,
			IdStatus: v.IdStatus,
			Caption:  v.Caption,
			Name:     v.Name,
			Role:     v.Role,
		})
	}
	return dataCmn
}

// SUBMISSION
func ToEntitySub(data Submission) mentee.Submission {
	return mentee.Submission{
		ID:        data.ID,
		ID_Mentee: data.IdMentee,
		Title:     data.Title,
		File:      data.File,
		Score:     data.Score,
		Status: data.Status,
	}
}

func FromEntitySub(data mentee.Submission) Submission {
	return Submission{
		Model:    gorm.Model{ID: data.ID},
		IdTask:   data.ID_Tasks,
		IdMentee: data.ID_Mentee,
		File:     data.File,
	}
}

// TASK
func ToCoreArrayTask(task []Task) []mentee.Task {
	var res []mentee.Task
	for _, val := range task {
		res = append(res, mentee.Task{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
			Images:      val.Images,
			DueDate:     *val.DueDate,
		})
	}
	return res
}
