package services

import (
	"be12/mentutor/features/login"
	"be12/mentutor/middlewares"
	"errors"
	"strings"
	"unicode"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authData login.DataInterface
}

func New(data login.DataInterface) login.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) Login(input login.Core) (login.Core, string, error) {

	// PENGECEKAN STRING EMAIL & PASSWORD
	
	// Cek email mengandung spasi
	for _, v := range input.Email{
		if unicode.IsSpace(v) {
			return login.Core{}, "", errors.New("contain space")
		}
	}
	// Cek ketentuan email
	if len(input.Email) < 10 || len(input.Email) > 75  {
		return login.Core{}, "", errors.New("length not valid")
	} else if strings.Contains(input.Email, "@") == false && strings.Contains(input.Email, ".") == false {
		return login.Core{}, "", errors.New("not contain (@) or (.)")
	} 

	// Cek ketentuan password
	var upper, lower, number, sChar int
	for _, v := range input.Password {
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

	if upper < 1 || lower < 1 || number < 1 || sChar < 1 {
		return login.Core{}, "", errors.New("string not as expected")
	} else if len(input.Password) < 8 || len(input.Password) > 30 {
		return login.Core{}, "", errors.New("string too short or too long")
	} 

	// Check Username di database
	res, err := usecase.authData.Login(input)
	if err != nil {
		log.Error(err.Error(), "username not found")
		return login.Core{}, "", errors.New("username not found")
	}
	//  Check password adalah sama
	if res.Role == "mentor" || res.Role == "mentee" {
		pass := login.Core{Password: res.Password}
		check := bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(input.Password))
		if check != nil {
			log.Error(check, " wrong password")
			return login.Core{}, "", errors.New("password not equal")
		}
	}
	token, err := middlewares.CreateToken(int(res.ID), int(res.IdClass), res.Role)

	return res, token, err
}
