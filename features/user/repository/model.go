package repository

import (
	"be12/mentutor/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Images   string
	Email    string
	Password string
	Role     string
	Class    string
}

func FromDomain(du user.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Password: du.Password,
		Images:   du.Images,
	}
}

func fromCore(dataCore user.Core) User {
	dataModel := User{
		// ID:     dataCore.ID,
		Name:     dataCore.Name,
		Images:   dataCore.Images,
		Class:    dataCore.Class,
		Role:     dataCore.Role,
		Email:    dataCore.Email,
		Password: dataCore.Password,
	}
	return dataModel

}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Images:   data.Images,
		Class:    data.Class,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}

}

func ToDomain(u User) user.Core {
	return user.Core{
		ID:       u.ID,
		Name:     u.Name,
		Class:    u.Class,
		Password: u.Password,
		Images:   u.Images,
		Role:     u.Role,
	}
}
func ToCoreList(data []User) []user.Core {
	var dataCore []user.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
