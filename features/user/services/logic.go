package services

import (
	"be12/mentutor/features/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}
func (usecase *userUsecase) InsertData(data user.Core) (int, error) {
	if data.Name == "" {
		return -1, errors.New("masukan nama ")
	}
	if data.Email == "" {
		return -1, errors.New("masukan email ")
	}
	if data.Password == "" {
		return -1, errors.New("masukann password")
	}
	if data.Role == "" {
		return -1, errors.New("masukann role")
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	data.Password = string(hashPass)
	row, err := usecase.userData.PostData(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}
