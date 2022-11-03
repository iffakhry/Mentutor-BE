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
	var mentee Mentee
	// Get data mentor
	err := repo.db.Where("email = ?", cnv.Email).First(&cnv).Error
	
	if err != nil {
		// Get data mentee
		if err := repo.db.Where("email = ?", cnv.Email).Model(&Mentee{}).Scan(&mentee).Error; err != nil {
			log.Print("error get data mentee")
			return login.Core{}, err
		}
	} 
	input = ToDomainMentor(cnv)
	return input, nil
	
}
