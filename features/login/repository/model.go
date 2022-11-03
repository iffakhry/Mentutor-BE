package repository

import (
	"be12/mentutor/features/login"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Images   string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:enum('admin','mentor');not null"`
	IdClass  uint
	Class string
}

func ToDomain(u Mentor) login.Core {
	return login.Core{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Role:     u.Role,
		IdClass: u.IdClass,
		Class: u.Class,
	}
}

func FromDomain(du login.Core) Mentor {
	return Mentor{
		Email:    du.Email,
		Password: du.Password,
	}
}
