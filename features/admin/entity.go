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
}

type UsecaseInterface interface {
	AddUser(input UserCore, c echo.Context) (UserCore, error)
}

type RepoInterface interface {
	InsertMentee(input UserCore) (UserCore, error)
	InsertMentor(input UserCore) (UserCore, error)
	GetAllClass(id uint) (ClassCore, error)
}

