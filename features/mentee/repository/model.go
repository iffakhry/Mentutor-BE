package repository

import (
	"be12/mentutor/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"type:varchar(255);not null;unique"`
	Password    string `gorm:"type:varchar(255);not null"`
	Images      string `gorm:"type:varchar(255);not null"`
	IdClass     uint
}

func FromEntity(data mentee.MenteeCore) Mentee {
	return Mentee{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		Images: data.Images,
	}
}

func ToEntity(id uint, data Mentee) mentee.MenteeCore {
	return mentee.MenteeCore{
		IdUser: id,
		Name: data.Name,
		Email: data.Email,
		Images: data.Images,
	}
}