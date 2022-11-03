package admin

import "github.com/labstack/echo/v4"

type UserCore struct {
	IdUser   uint
	Name     string
	Email    string
	IdClass  uint
	Class    string
	Password string
	Role     string
}

type ClassCore struct {
	IdClass uint
	ClassName string
	Status string
	TotalStudent int
}

type UsecaseInterface interface {
	AddUser(input UserCore, c echo.Context) (UserCore, error)
	GetAllUser( c echo.Context) ([]UserCore, []UserCore, error)
	AddNewClass(input ClassCore, c echo.Context) (error)
	GetAllClass(c echo.Context) ([]ClassCore, error)
}

type RepoInterface interface {
	InsertMentee(input UserCore) (UserCore, error)
	InsertMentor(input UserCore) (UserCore, error)
	GetClass(id uint) (ClassCore, error)
	GetAllUser() ([]UserCore, []UserCore,error)
	InsertNewClass(input ClassCore) (error)
	GetAllClass() ([]ClassCore, error)
}

