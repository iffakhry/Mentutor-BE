package repository

import (
	"be12/mentutor/features/mentee"
	"log"

	"gorm.io/gorm"
)

type menteeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.RepoInterface {
	return &menteeData{
		db: db,
	}
}

func (md *menteeData) EditProfile(id uint, data mentee.MenteeCore) (mentee.MenteeCore, error) {
	var res Mentee

	res = FromEntity(data)
	if err := md.db.Where("id = ?", id).Updates(&res).Error; err != nil {
		log.Print(err.Error(), "ERROR INSERT TO DATBASE")
		return mentee.MenteeCore{}, err
	}

	cnv := ToEntity(id, res)

	return cnv, nil
}

func (md *menteeData) AddStatus(data mentee.Status, token int) (mentee.Status, error) {

	dataModel := ToEntityMentee(data)
	dataModel.IdMentee = uint(token)
	tx := md.db.Create(&dataModel).Last(&dataModel)
	log.Print(dataModel.ID, "ini id status")
	if tx.Error != nil {
		return mentee.Status{}, tx.Error
	}
	AddRes := toPostUser(dataModel)
	AddRes.ID = dataModel.ID

	return AddRes, nil

}

func (md *menteeData) GetAllPosts() ([]mentee.Status, error) {
	var dataPost []Status
	tx := md.db.Find(&dataPost)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataPostUser := ToCoreArray(dataPost)

	return dataPostUser, nil
}
