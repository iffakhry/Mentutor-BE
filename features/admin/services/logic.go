package services

import (
	"be12/mentutor/features/admin"
	"errors"
	"log"
	"strconv"

	// "log"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepo admin.RepoInterface
}

func New(data admin.RepoInterface) admin.UsecaseInterface {
	return &adminUsecase{
		adminRepo: data,
	}
}

func (au *adminUsecase) AddUser(input admin.UserCore, role string) (admin.UserCore, error) {

	// CEK ROLE USE(c)
	if role != "admin" {
		log.Print("user not admin")
		return admin.UserCore{}, errors.New("user not admin")
	}

	// CEK KONDISI EMAIL
	for _, v := range input.Email {
		if unicode.IsSpace(v) {
			log.Print("contain space")
			return admin.UserCore{}, errors.New("contain space")
		}
	}
	if len(input.Email) < 8 || len(input.Email) > 40 {
		return admin.UserCore{}, errors.New("length email not valid")
	} else if strings.Contains(input.Email, "@") == false || strings.Contains(input.Email, ".") == false {
		return admin.UserCore{}, errors.New("not contain (@) and (.)")
	}

	// String to lower email
	strEmail := strings.ToLower(input.Email)
	input.Email = strEmail

	// CEK KONDISI NAMA
	if len(input.Name) < 5 || len(input.Name) > 50 {
		return admin.UserCore{}, errors.New("length name not valid")
	}
	var upper, lower, number, sChar int
	for _, v := range input.Name {
		if unicode.IsUpper(v) == true {
			upper += 1
		} else if unicode.IsLower(v) == true {
			lower += 1
		} else if unicode.IsNumber(v) == true {
			number+=1
		} else if unicode.IsPunct(v) {
			sChar+=1
		}
	}

	if upper < 1 {
		return admin.UserCore{}, errors.New("Ainput name not valid")
	} else if lower < 1 {
		return admin.UserCore{}, errors.New("Binput name not valid")
	} else if number > 1 {
		return admin.UserCore{}, errors.New("Cinput name not valid")
	} else if sChar > 1 {
		return admin.UserCore{}, errors.New("Dinput name not valid")
	}

	// CEK KONDISI PASSOWRD
	var sCharString = "@#$%^&*<>:;'[]{}|`~!"
	var passUpper, passLower, passNumber, specialChar int
	for _, v := range input.Password {
		if unicode.IsUpper(v) == true {
			passUpper += 1
		} else if unicode.IsLower(v) == true {
			passLower += 1
		} else if unicode.IsNumber(v) == true {
			passNumber += 1
		} else if unicode.IsPunct(v) == true {
			specialChar += 1
		} else if strings.Contains(sCharString, string(v)) == true {
			specialChar += 1
		}
	}
	if passUpper < 1 {
		return admin.UserCore{}, errors.New("string not as expected")
	} else if passLower < 1 {
		return admin.UserCore{}, errors.New("string not as expected")
	} else if passNumber < 1 {
		return admin.UserCore{}, errors.New("string not as expected")
	}
	if specialChar == 0 {
		return admin.UserCore{}, errors.New("string not as expected")
	} else if len(input.Password) < 8 || len(input.Password) > 30 {
		return admin.UserCore{}, errors.New("string too short or too long")
	}

	// CEK KELAS TERSEDIA
	idClass := uint(input.IdClass)
	_, err := au.adminRepo.GetClass(idClass)
	if err != nil {
		return admin.UserCore{}, errors.New("input class not valid")
	}

	// ENKRIPSI PASSWORD
	generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(generate)

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
			log.Print("error add mentor")
			return admin.UserCore{}, errors.New("error add user")
		} else {
			return res, nil
		}
	}
}

func (au *adminUsecase) GetAllUser(role string) ([]admin.UserCore, []admin.UserCore, error) {

	if role != "admin" {
		return []admin.UserCore{}, []admin.UserCore{}, errors.New("user not admin")
	}

	resMentee, resMentor, err := au.adminRepo.GetAllUser()
	if err != nil {
		return []admin.UserCore{}, []admin.UserCore{}, errors.New("query error")
	}
	return resMentee, resMentor, nil
}

func (au *adminUsecase) AddNewClass(input admin.ClassCore, role string) (admin.ClassCore, error) {

	if role != "admin" {
		return admin.ClassCore{}, errors.New("user not admin")
	}
	// CEK KONDISI NAMA KELAS
	if len(input.ClassName) < 5 || len(input.ClassName) > 30 {
		return admin.ClassCore{}, errors.New("length name not valid")
	}

	var sChar int
	for _, v := range input.ClassName {
		if unicode.IsPunct(v) == true{
			sChar+=1
		} else if string(v) == "?" {
			sChar+=1
		}
		_, err := strconv.Atoi(input.ClassName)
		if err == nil {
			log.Print(err)
			return admin.ClassCore{}, errors.New("class name is number")
		}
	}

	if sChar > 1 {
		return admin.ClassCore{}, errors.New("contain special character")
	}

	res, err := au.adminRepo.InsertNewClass(input)
	if err != nil {
		return admin.ClassCore{}, errors.New("input not valid")
	} else if res.IdClass == 0 {
		return admin.ClassCore{}, errors.New("input not valid")
	}
	return res, nil
}

func (au *adminUsecase) GetAllClass(role string) ([]admin.ClassCore, error) {

	if role != "admin" {
		return []admin.ClassCore{}, errors.New("user not admin")
	}

	res, err := au.adminRepo.GetAllClass()
	if err != nil {
		return []admin.ClassCore{}, errors.New("error in database")
	}

	return res, nil
}

func (au *adminUsecase) UpdateUserAdmin(input admin.UserCore, role string) (admin.UserCore, error) {

	if role != "admin" {
		return admin.UserCore{}, errors.New("user not admin")
	}

	var user admin.UserCore

	// CEK SEMUA DATA KOSONG
	if input.Name == "" && input.Email == "" && input.IdClass == 0 && input.Password == "" && input.Images == "" {
		return admin.UserCore{}, errors.New("no data input")
	}

	// AMBIL DATA DARI DATABASE
	if input.IdUser < 1000 {
		res, err := au.adminRepo.GetSingleMentee(input.IdUser)
		if err != nil {
			return admin.UserCore{}, errors.New("user not found")
		} else if res.IdUser == 0 {
			return admin.UserCore{}, errors.New("user not found")
		}
		user = res
	} else if input.IdUser >= 1000 {
		res, err := au.adminRepo.GetSingleMentor(input.IdUser)
		if err != nil {
			return admin.UserCore{}, errors.New("user not found")
		} else if res.IdUser == 0 {
			return admin.UserCore{}, errors.New("user not found")
		}
		user = res
	}

	// CEK KONDISI NAMA
	if input.Name != "" {
		if len(input.Name) < 5 || len(input.Name) > 50 {
			return admin.UserCore{}, errors.New("length name not valid")
		}
		var upper, lower, number, sChar, space int
		for _, v := range input.Name {
			if unicode.IsUpper(v) == true {
				upper += 1
			} else if unicode.IsLower(v) == true {
				lower += 1
			} else if unicode.IsNumber(v) == true {
				number += 1
			} else if unicode.IsPunct(v) {
				sChar += 1
			} else if unicode.IsSpace(v) {
				space += 1
			}
		}
		if upper < 1 {
			return admin.UserCore{}, errors.New("input name not valid")
		} else if lower < 1 {
			return admin.UserCore{}, errors.New("input name not valid")
		} else if number > 0 {
			return admin.UserCore{}, errors.New("input name not valid")
		} else if sChar > 0 {
			return admin.UserCore{}, errors.New("input name not valid")
		} else if space < 1 {
			return admin.UserCore{}, errors.New("input name not valid")
		}
	} else {
		input.Name = user.Name
	}

	// CEK KONDISI EMAIL
	if input.Email != "" {
		for _, v := range input.Email {
			if unicode.IsSpace(v) {
				log.Print("contain space")
				return admin.UserCore{}, errors.New("contain space")
			}
		}
		contain1 := strings.Contains(input.Email, "@") 
		contain2 := strings.Contains(input.Email, ".") 
		if len(input.Email) < 8 || len(input.Email) > 40 {
			return admin.UserCore{}, errors.New("length email not valid")
		} else if contain1 == false || contain2 == false {
			return admin.UserCore{}, errors.New("not contain (@) and (.)")
		}
		tmp := strings.ToLower(input.Email)
		input.Email = tmp

	} else {
		strings.ToLower(user.Email)
		input.Email = user.Email
	}

	

	// CEK KONDISI PASSWORD
	if input.Password != "" {
		var sChar = "@#$%^&*<>:;'[]{}|`~!"
		var passUpper, passLower, passNumber, specialChar int
		for _, v := range input.Password {
			if unicode.IsUpper(v) == true {
				passUpper += 1
			} else if unicode.IsLower(v) == true {
				passLower += 1
			} else if unicode.IsNumber(v) == true {
				passNumber += 1
			} else if unicode.IsPunct(v) == true {
				specialChar += 1
			} else if strings.Contains(sChar, string(v)) == true {
				specialChar += 1
			}
		}
		if passUpper < 1 {
			return admin.UserCore{}, errors.New("string not as expected")
		} else if passLower < 1 {
			return admin.UserCore{}, errors.New("string not as expected")
		} else if passNumber < 1 {
			return admin.UserCore{}, errors.New("string not as expected")
		}
		if specialChar == 0 {
			return admin.UserCore{}, errors.New("string not as expected")
		} else if len(input.Password) < 8 || len(input.Password) > 30 {
			return admin.UserCore{}, errors.New("string too short or too long")
		}

		generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		input.Password = string(generate)
	}

	// CEK KELAS TERSEDIA]
	if input.IdClass == 0 {
		input.IdClass = user.IdClass
	} else if input.IdClass != 0 {
		idClass := uint(input.IdClass)
		_, err := au.adminRepo.GetClass(idClass)
		if err != nil {
			return admin.UserCore{}, errors.New("input class not valid")
		}
	}

	if input.IdUser >= 1000 {
		res, err := au.adminRepo.EditUserMentor(input)
		if err != nil {
			return admin.UserCore{}, errors.New("error in database")
		}

		return res, nil
	} else if input.IdUser < 1000 {
		res, err := au.adminRepo.EditUserMentee(input)
		if err != nil {
			input.IdUser = 0
			return admin.UserCore{}, errors.New("error in database")
		} else {
			return res, nil
		}
	}
	return admin.UserCore{}, errors.New("user not found")
}

func (au *adminUsecase) DeleteUser(id uint, role string) error {

	if role != "admin" {
		return errors.New("user not admin")
	}

	if id < 1000 {
		err := au.adminRepo.DeleteUserMentee(id)
		if err != nil {
			log.Print("eror in database")
			return errors.New("error in database")
		}
		return nil
	} else if id >= 1000 {
		err := au.adminRepo.DeleteUserMentor(id)
		if err != nil {
			return errors.New("error in database")
		}
		return nil
	}
	return errors.New("error in database")
}

func (au *adminUsecase) GetSingleUser(id uint, role string) (admin.UserCore, error) {

	if role != "admin" || role != "mentee" || role != "mentor"{
		return admin.UserCore{}, errors.New("user not admin")
	}

	if id < 1000 {
		res, err := au.adminRepo.GetSingleMentee(id)
		if err != nil {
			return admin.UserCore{}, errors.New("error in database")
		}
		if res.IdUser == 0 {
			return admin.UserCore{}, errors.New("error in database")
		}
		return res, nil
	} else if id >= 1000 {
		res, err := au.adminRepo.GetSingleMentor(id)
		if err != nil {
			return admin.UserCore{}, errors.New("error in database")
		}
		if res.IdUser == 0 {
			return admin.UserCore{}, errors.New("error in database")
		}
		return res, nil
	}
	return admin.UserCore{}, errors.New("error in database")
}

func (au *adminUsecase) UpdateClass(input admin.ClassCore, role string) (admin.ClassCore, error) {

	if role != "admin" {
		return admin.ClassCore{}, errors.New("user not admin")
	}

	// CEK KONDISI NAMA KELAS
	if len(input.ClassName) < 5 == true || len(input.ClassName) > 30 == true {
		return admin.ClassCore{}, errors.New("length name not valid")
	}

	var sChar int
	for _, v := range input.ClassName {
		if unicode.IsPunct(v) == true{
			sChar+=1
		} else if string(v) == "?" {
			sChar+=1
		}
		_, err := strconv.Atoi(input.ClassName)
		if err == nil {
			return admin.ClassCore{}, errors.New("class name is number")
		}
	}

	if sChar > 1 {
		return admin.ClassCore{}, errors.New("contain special character")
	}

	// CEK KELAS TERSEDIA
	dataClass, err := au.adminRepo.GetSingleClass(input.IdClass)
	if err != nil {
		return admin.ClassCore{}, err
	}

	if input.Status == "" {
		input.Status = dataClass.Status
	}

	if input.ClassName == "" {
		input.ClassName = dataClass.ClassName
	}

	res, err := au.adminRepo.EditClass(input)
	if err != nil {
		return admin.ClassCore{}, errors.New("error in database")
	}
	return res, nil
}

func (au *adminUsecase) DeleteClass(id uint, role string) error {

	if role != "admin" {
		return errors.New("user not admin")
	}

	err := au.adminRepo.DeleteClass(id)
	if err != nil {
		return errors.New("error in databse")
	}
	return nil
}
