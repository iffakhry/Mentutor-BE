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

func (ar *adminRepo) GetClass(id uint) (admin.ClassCore, error) {
	var res Class
	
	if err := ar.db.Where("status = ? AND id = ?", "active", id).First(&res).Error; err != nil {
		return admin.ClassCore{}, err
	}
	cnv := ToDomainClass(res)
	return cnv, nil
}

func (ar *adminRepo) GetAllUser() ([]admin.UserCore, []admin.UserCore, error) {
	var mentees []Mentee
	var mentors []Mentor
	
	if err := ar.db.Find(&mentees).Error; err != nil {
		return []admin.UserCore{}, []admin.UserCore{}, err
	}

	if err := ar.db.Where("role != ?", "admin").Find(&mentors).Error; err != nil {
		return []admin.UserCore{}, []admin.UserCore{}, err
	}

	cnvMentees := ToDomainMenteeArray(mentees)
	cnvMentors := ToDomainMentorArray(mentors)

	return cnvMentees, cnvMentors, nil

}

func (ar *adminRepo) InsertNewClass(input admin.ClassCore) error {
	data := FromDomainClass(input)
	if err := ar.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepo) GetAllClass() ([]admin.ClassCore, error) {
	var res []GetClass
	

	if err := ar.db.Model(&Class{}).Select("id, class_name, status").Scan(&res).Error; err != nil {
		return []admin.ClassCore{}, err
	}

	for i, val := range res{
		var count int64
		var tmp Mentee
		if err := ar.db.Model(&Mentee{}).Where("id_class = ?", val.ID).Find(&tmp).Count(&count).Error; err != nil {
			return []admin.ClassCore{}, err
		}
		res[i].TotalStudent = int(count)
	}
	
	cnv := ToDomainClassArray(res)
	return cnv, nil
}