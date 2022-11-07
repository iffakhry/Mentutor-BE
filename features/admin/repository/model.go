package repository

import (
	"be12/mentutor/features/admin"
	"time"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Images    string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:enum('mentee');not null"`
	ClassName string
	IdClass   uint
}

type Mentor struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Images    string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:enum('admin','mentor');not null"`
	ClassName string
	IdClass   uint
}

type MenteeSingle struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Images    string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:enum('mentee');not null"`
	ClassName string
	IdClass   uint
}

type MentorSingle struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Images    string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:enum('admin','mentor');not null"`
	ClassName string
	IdClass   uint
}

type Class struct {
	gorm.Model
	ClassName string
	Status    string
}

type GetClass struct {
	gorm.Model
	ClassName    string
	Status       string
	TotalStudent int
}

type Task struct {
	gorm.Model
	IdClass     uint
	IdMentor    uint
	Title       string
	Description string
	File        string
	Images      string
	DueDate     *time.Time
}

type Submission struct {
	gorm.Model
	IdMentee uint
	IdTask   uint
	File     string
	Score    int
}

type Status struct {
	gorm.Model
	IdMentee uint
	Caption  string
	Images   string
}

type Comment struct {
	gorm.Model
	IdUser   uint
	IdStatus uint
	Caption  string
}

// FROM DOMAIN

func FromDomainMentee(data admin.UserCore) Mentee {
	return Mentee{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		IdClass:  data.IdClass,
		Role:     data.Role,
	}
}

func FromDomainMentor(data admin.UserCore) Mentor {
	return Mentor{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		IdClass:  data.IdClass,
		Role:     data.Role,
	}
}

func FromDomainClass(data admin.ClassCore) Class {
	return Class{
		ClassName: data.ClassName,
		Status:    "active",
	}
}

func FromDomainUpdateClass(data admin.ClassCore) Class {
	return Class{
		Model:     gorm.Model{ID: data.IdClass},
		ClassName: data.ClassName,
		Status:    data.Status,
	}
}

func FromDomainUpdateMentee(data admin.UserCore) Mentee {
	return Mentee{
		Model:    gorm.Model{ID: data.IdUser},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Images:   data.Images,
		Role:     data.Role,
		IdClass:  data.IdClass,
	}
}

func FromDomainUpdateMentor(data admin.UserCore) Mentor {
	return Mentor{
		Model:    gorm.Model{ID: data.IdUser},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Images:   data.Images,
		Role:     data.Role,
		IdClass:  data.IdClass,
	}
}

// TO DOMAIN

func ToDomainMentee(data Mentee) admin.UserCore {
	return admin.UserCore{
		IdUser:  data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Images:  data.Images,
		IdClass: data.IdClass,
		Role:    data.Role,
	}
}

func ToDomainMentor(data Mentor) admin.UserCore {
	return admin.UserCore{
		IdUser:  data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Images:  data.Images,
		IdClass: data.IdClass,
		Role:    data.Role,
	}
}

func ToDomainSingleMentee(data MenteeSingle) admin.UserCore {
	return admin.UserCore{
		IdUser:  data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Images:  data.Images,
		IdClass: data.IdClass,
		Class:   data.ClassName,
		Role:    data.Role,
	}
}

func ToDomainSingleMentor(data MentorSingle) admin.UserCore {
	return admin.UserCore{
		IdUser:  data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Images:  data.Images,
		IdClass: data.IdClass,
		Class:   data.ClassName,
		Role:    data.Role,
	}
}

func ToDomainClass(data Class) admin.ClassCore {
	return admin.ClassCore{
		IdClass:   data.ID,
		ClassName: data.ClassName,
		Status:    data.Status,
	}
}

func ToDomainClassArray(data []GetClass) []admin.ClassCore {
	var res []admin.ClassCore

	for _, val := range data {
		res = append(res, admin.ClassCore{
			IdClass:      val.ID,
			ClassName:    val.ClassName,
			Status:       val.Status,
			TotalStudent: val.TotalStudent,
		})
	}
	return res
}

func ToDomainMenteeArray(data []Mentee) []admin.UserCore {
	var res []admin.UserCore

	for _, val := range data {
		res = append(res, admin.UserCore{
			IdUser: val.ID,
			Name:   val.Name,
			Email:  val.Email,
			Role:   val.Role,
			Class:  val.ClassName,
		})
	}
	return res
}

func ToDomainMentorArray(data []Mentor) []admin.UserCore {
	var res []admin.UserCore

	for _, val := range data {
		res = append(res, admin.UserCore{
			IdUser: val.ID,
			Name:   val.Name,
			Email:  val.Email,
			Role:   val.Role,
			Class:  val.ClassName,
		})
	}
	return res
}
