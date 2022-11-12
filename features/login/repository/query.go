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
	// cnvMentee := FromDomainMentee(input)
	var mentee Mentee
	// Get data mentor
	err := repo.db.Where("email = ?", cnv.Email).First(&cnv).Error
	
	if err != nil {
		// Get data mentee
		if err := repo.db.Where("email = ?", cnv.Email).Model(&Mentee{}).Scan(&mentee).Error; err != nil {
			log.Print("error get data mentee")
			return login.Core{}, err
		}
		repo.db.Model(&Class{}).Where("id = ?", mentee.IdClass).Select("classes.class_name").Scan(&mentee)
		input = ToDomainMentee(mentee)
		return input, nil
	} 
	repo.db.Model(&Class{}).Where("id = ?", cnv.IdClass).Select("classes.class_name").Scan(&cnv)
	input = ToDomainMentor(cnv)
	return input, nil
}

func (ad *authData) GetToken(id uint) error {

	if err := ad.db.Where("id_mentee = ?", id).First(&GmailToken{}).Error; err != nil {
		return err
	}
	return nil
}

func (ad *authData) InsertToken(idMentee uint) error {
	var token GmailToken

	token.IdMentee = idMentee

	if err := ad.db.Create(&token).Error; err != nil {
		return err
	}
	return nil
}