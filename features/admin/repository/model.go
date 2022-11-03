package repository

import (
	"be12/mentutor/features/admin"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"type:varchar(255);not null;unique"`
	Password    string `gorm:"type:varchar(255);not null"`
	Images      string `gorm:"type:varchar(255);not null"`
	Role        string `gorm:"type:enum('mentee');not null"`
	IdClass     uint
}

type Mentor struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Images   string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:enum('admin','mentor');not null"`
	IdClass  uint
}

type Class struct {
	gorm.Model
	ClassName string 
	Status    string 
	IdMentor  uint
}

func FromDomainMentee(data admin.UserCore) Mentee {
	return Mentee{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		IdClass: data.IdClass,
		Role: data.Role,
	}
}

func FromDomainMentor(data admin.UserCore) Mentor {
	return Mentor{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		IdClass: data.IdClass,
		Role: data.Role,
	}
}

func ToDomainMentee(data Mentee) admin.UserCore {
	return admin.UserCore{
		IdUser: data.ID,
		Name: data.Name,
		Email: data.Email,
		IdClass: data.IdClass,
		Role: data.Role,
	}
}

func ToDomainMentor(data Mentor) admin.UserCore {
	return admin.UserCore{
		IdUser: data.ID,
		Name: data.Name,
		Email: data.Email,
		IdClass: data.IdClass,
		Role: data.Role,
	}
}

func ToDomainClass(data Class) admin.ClassCore {
	return admin.ClassCore{
		IdClass: data.ID,
		ClassName: data.ClassName,
	}
}

func ToDomainMenteeArray(data []Mentee) []admin.UserCore{
	var res []admin.UserCore
	
	for _, val := range data {
		res = append(res, admin.UserCore{
			IdUser: val.ID,
			Name: val.Name,
			Role: val.Role,
			Class: val.Role,
		})
	}
	return res
}

func ToDomainMentorArray(data []Mentor) []admin.UserCore{
	var res []admin.UserCore
	
	for _, val := range data {
		res = append(res, admin.UserCore{
			IdUser: val.ID,
			Name: val.Name,
			Role: val.Role,
			Class: val.Role,
		})
	}
	return res
}