package repository

import (
	"be12/mentutor/features/login"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Class    string `json:"class" form:"class"`
}

func ToDomain(u Mentor) login.Core {
	return login.Core{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Role:     u.Role,
		Class:    u.Class,
	}
}

func FromDomain(du login.Core) Mentor {
	return Mentor{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Password: du.Password,
	}
}
