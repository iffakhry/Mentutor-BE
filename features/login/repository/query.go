package repository

import (
	"be12/mentutor/features/login"
	"log"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) login.DataInterface {
	return &authData{
		db: db,
	}
}
func (repo *authData) Login(input login.Core) (login.Core, error) {

	cnv := FromDomain(input)
	if err := repo.db.Where("email = ?", cnv.Email).First(&cnv).Error; err != nil {
		log.Print(" error get data")
		return login.Core{}, err
	}
	input = ToDomain(cnv)
	return input, nil
}
