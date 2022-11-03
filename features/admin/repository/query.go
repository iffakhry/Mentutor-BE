package repository

import (
	"be12/mentutor/features/admin"

	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) admin.RepoInterface {
	return &adminRepo{
		db: db,
	}
}

func (ar *adminRepo) InsertMentee(data admin.UserCore) (admin.UserCore, error) {
	input := FromDomainMentee(data)
	if err := ar.db.Create(&input).Error; err != nil {
		return admin.UserCore{}, err
	}

	return data, nil  
}

func (ar *adminRepo) InsertMentor(data admin.UserCore) (admin.UserCore, error) {
	input := FromDomainMentor(data)
	if err := ar.db.Create(&input).Error; err != nil {
		return admin.UserCore{}, err
	}
	return data, nil  
}

func (ar *adminRepo) GetAllClass(id uint) (admin.ClassCore, error) {
	var res Class
	
	if err := ar.db.Where("status = ?", "aktif").First(&res).Error; err != nil {
		return admin.ClassCore{}, err
	}
	cnv := ToDomainClass(res)
	return cnv, nil
}