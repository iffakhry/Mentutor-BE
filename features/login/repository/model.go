package repository

import (
	"be12/mentutor/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Class    string `json:"class" form:"class"`
}

func ToDomain(u User) login.Core {
	return login.Core{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Role:     u.Role,
		Class:    u.Class,
	}
}

func FromDomain(du login.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Password: du.Password,
	}
}
