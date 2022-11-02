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

func (service *userUsecase) GetProfile(token int) (user.Core, error) {
	dataId, err := service.userData.MyProfile(token)
	if err != nil {
		return user.Core{}, err
	}
	return dataId, nil
}
func (usecase *userUsecase) PutDataId(data user.Core) (int, error) {
	// row, err := usecase.userData.UpdateData(data)
	// return row, err
	if data.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(hash)
	}

	if data.Role == "admin" {
		data.Role = "Admin"
	}

	row, _ := usecase.userData.UpdateData(data)
	if row == -1 {
		return -1, errors.New("not found")
	}

	return row, nil
}
func (service *userUsecase) GetAlluser() ([]user.Core, error) {
	row, err := service.userData.GetAll()
	if err != nil {
		return []user.Core{}, err
	}
	return row, err
}

func (service *userUsecase) GetDataId(param, token int) (user.Core, error) {
	dataId, err := service.userData.SelectDataId(param, token)
	if err != nil {
		return user.Core{}, err
	}

	return dataId, nil
}

func (service *userUsecase) Delete(token int) (int, error) {
	row, err := service.userData.DeleteData(token)
	if err != nil {
		return -1, err
	}
	return row, nil
}
