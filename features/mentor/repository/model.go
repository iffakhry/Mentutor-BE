package repository

import (
	"be12/mentutor/features/mentor"
	"time"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Images    string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:enum('admin','mentor');not null"`
	ClassName string
	IdClass   uint
}

type Mentee struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Images    string `gorm:"type:varchar(255)"`
	Role      string `gorm:"type:enum('mentee');not null"`
	ClassName string
	IdClass   uint
}

type Class struct {
	gorm.Model
	ClassName string   `gorm:"type:varchar(255);unique;not null"`
	Status    string   `gorm:"type:enum('active','non_active');not null"`
	Mentors   []Mentor `gorm:"foreignKey:IdClass"`
}

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	IdClass     uint
	IdMentor    uint
	Description string `gorm:"type:varchar(255);not null"`
	File        string `gorm:"type:varchar(255);not null"`
	Images      string `gorm:"type:varchar(255);not null"`
	DueDate     time.Time
}

type Comment struct {
	gorm.Model
	IdUser   uint
	IdStatus uint
	Caption  string `gorm:"type:varchar(255);not null"`
}

type Submission struct {
	gorm.Model
	IdMentee uint
	IdTask   uint
	Name     string	`gorm:"<-:false"`
	Title    string	`gorm:"<-:false"`
	File     string `json:"file"`
	Score    int    `gorm:"type:int(3);not null"`
}

// FROM DOMAIN

func FromDomainMentee(data mentor.UserCore) Mentee {
	return Mentee{
		Model:    gorm.Model{ID: data.IdUser},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Images:   data.Images,
		IdClass:  data.IdClass,
	}
}

func FromDomainMentor(data mentor.UserCore) Mentor {
	return Mentor{
		Model:    gorm.Model{ID: data.IdUser},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Images:   data.Images,
		IdClass:  data.IdClass,
	}
}

func FromDomainTask(data mentor.TaskCore) Task {
	return Task{
		IdClass:     data.IdClass,
		IdMentor:    data.IdMentor,
		Description: data.Description,
		File:        data.File,
		Images:      data.Images,
		DueDate:     data.DueDate,
		Title:       data.Title,
	}
}

func FromDomainSubmission(data mentor.SubmissionCore) Submission {
	return Submission{
		Model:  gorm.Model{ID: data.ID},
		IdTask: data.IdTask,
		Score:  data.Score,
	}
}

// TO DOMAIN

func ToDomainMentee(data Mentee) mentor.UserCore {
	return mentor.UserCore{
		IdUser: data.ID,
		Name:   data.Name,
		Email:  data.Email,
		Images: data.Images,
	}
}

func ToDomainMentor(data Mentor) mentor.UserCore {
	return mentor.UserCore{
		IdUser: data.ID,
		Name:   data.Name,
		Email:  data.Email,
		Images: data.Images,
	}
}

func ToDomainTask(data Task) mentor.TaskCore {
	return mentor.TaskCore{
		ID:          data.ID,
		IdClass:     data.IdClass,
		IdMentor:    data.IdMentor,
		Title:       data.Title,
		Description: data.Description,
		File:        data.File,
		Images:      data.Images,
		DueDate:     data.DueDate,
	}
}

func ToDomainAllTask(data []Task) []mentor.TaskCore {
	var res []mentor.TaskCore

	for _, val := range data {
		res = append(res, mentor.TaskCore{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
			Images:      val.Images,
			File:        val.File,
			DueDate:     val.DueDate,
		})
	}
	return res
}

func ToDomainSingleTask(data Task) mentor.TaskCore {
	return mentor.TaskCore{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Images:      data.Images,
		File:        data.File,
		DueDate:     data.DueDate,
	}
}

func ToDomainTaskSub(data []Submission) []mentor.SubmissionCore {
	var res []mentor.SubmissionCore

	for _, val := range data {
		res = append(res, mentor.SubmissionCore{
			ID:         val.ID,
			File:       val.File,
			IdTask:     val.IdTask,
			Score:      val.Score,
			NameMentee: val.Name,
		})
	}
	return res
}

func ToDomainSubmission(data Submission) mentor.SubmissionCore {
	return mentor.SubmissionCore{
		ID:    data.ID,
		Title: data.Title,
		File: data.File,
		Score: data.Score,
	}
}
