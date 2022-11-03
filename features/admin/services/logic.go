package services

import (
	"be12/mentutor/features/admin"
	"be12/mentutor/middlewares"
	"errors"
	"strings"
	"unicode"

	"github.com/labstack/echo/v4"
)

type adminUsecase struct {
	adminRepo admin.RepoInterface
}

func New(data admin.RepoInterface) admin.UsecaseInterface {
	return &adminUsecase{
		adminRepo: data,
	}
}

func (au *adminUsecase) AddUser(input admin.UserCore, c echo.Context) (admin.UserCore, error) {

	// CEK ROLE USER
	_, _, role := middlewares.ExtractToken(c)
	if role != "admin" {
		return admin.UserCore{}, errors.New("user not admin")
	}

	// CEK KONDISI EMAIL
	if len(input.Email) < 8 || len(input.Email) > 40  {
		return admin.UserCore{}, errors.New("input not valid")
	} else if strings.Contains(input.Email, "@") == false && strings.Contains(input.Email, ".") == false {
		return admin.UserCore{}, errors.New("input not valid")
	} 

	// CEK KONDISI NAMA
	if len(input.Name) < 5 || len(input.Name) > 50{
		return admin.UserCore{}, errors.New("input not valid")
	}
	var upper, lower, number, sChar int
	for _, v := range input.Name {
		if unicode.IsUpper(v) == true {
			upper+=1
		} else if unicode.IsLower(v) ==  true {
			lower += 1
		} else if unicode.IsNumber(v) == true {
			number+=1
		} else{
			sChar+=1
		}
	}

	if upper < 1 {
		return admin.UserCore{}, errors.New("input name not valid")
	} else if lower < 1 {
		return admin.UserCore{}, errors.New("input not valid")
	} else if number > 1 {
		return admin.UserCore{}, errors.New("input not valid")
	} else if sChar > 1 {
		return admin.UserCore{}, errors.New("input not valid")
	}


	// CEK KELAS TERSEDIA
	idClass := uint(input.IdClass)
	_, err := au.adminRepo.GetClass(idClass)
	if err != nil {
		return admin.UserCore{}, errors.New("input not valid")
	}


	if input.Role == "mentee" {
		res, err := au.adminRepo.InsertMentee(input)
		if err != nil {
			return admin.UserCore{}, errors.New("error add user")
		} else {
			return res, nil
		}
	} else {
		res, err := au.adminRepo.InsertMentor(input)
		if err != nil {
			return admin.UserCore{}, errors.New("error add user")
		} else {
			return res, nil
		}
	}
}

func (au *adminUsecase) GetAllUser(c echo.Context) ([]admin.UserCore, []admin.UserCore, error) {
	_, _, role := middlewares.ExtractToken(c)
	if role != "admin" {
		return []admin.UserCore{}, []admin.UserCore{},errors.New("user not admin")
	}

	resMentee, resMentor, err := au.adminRepo.GetAllUser()
	if err != nil {
		return []admin.UserCore{}, []admin.UserCore{}, errors.New("query error")
	}
	return resMentee, resMentor, nil
}

func (au *adminUsecase) AddNewClass(input admin.ClassCore, c echo.Context) error {
	_, _, role := middlewares.ExtractToken(c)
	if role != "admin" {
		return errors.New("user not admin")
	}

	err := au.adminRepo.InsertNewClass(input)
	if err != nil {
		return errors.New("input not valid")
	}
	return nil
}

func (au *adminUsecase) GetAllClass(c echo.Context) ([]admin.ClassCore, error) {
	_, _, role := middlewares.ExtractToken(c)
	if role != "admin" {
		return []admin.ClassCore{}, errors.New("user not admin")
	}

	res, err := au.adminRepo.GetAllClass()
	if err != nil {
		return []admin.ClassCore{}, errors.New("error in database")
	}

	return res, nil
}