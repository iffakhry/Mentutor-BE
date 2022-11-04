package repository

import (
	"be12/mentutor/features/admin"
	"log"

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

func (ar *adminRepo) EditUserMentee(input admin.UserCore) (admin.UserCore, error) {
	data := FromDomainUpdateMentee(input)

	log.Print(data.ID)

	if err := ar.db.Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return admin.UserCore{}, err
	}
	return input, nil
}

func (ar *adminRepo) EditUserMentor(input admin.UserCore) (admin.UserCore, error) {
	data := FromDomainUpdateMentor(input)
	
	if err := ar.db.Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return admin.UserCore{}, err
	}
	return input, nil
}

func (ar *adminRepo) DeleteUserMentor(id uint) (error) {
	var mentor Mentor

	if err := ar.db.Where("id = ?", id).Delete(&mentor).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepo) DeleteUserMentee(id uint) (error) {
	var mentee Mentee
	mentee.ID = id

	if err := ar.db.Delete(&mentee).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepo) GetSingleMentee(id uint) (admin.UserCore, error) {

	var mentee MenteeSingle
	mentee.ID = id

	if err := ar.db.Model(&Mentee{}).
	Select("mentees.id, mentees.id_class, mentees.role, mentees.email, mentees.name, mentees.images, classes.class_name").
	Joins("left join classes on classes.id = mentees.id_class").
	Where("mentees.id = ?", id).Scan(&mentee).Error; err != nil {
		return admin.UserCore{}, err
	}
	cnv := ToDomainSingleMentee(mentee)
	return cnv, nil
} 

func (ar *adminRepo) GetSingleMentor(id uint) (admin.UserCore, error) {
	var mentor MentorSingle
	mentor.ID = id

	if err := ar.db.Model(&Mentor{}).
	Select("mentors.id, mentors.id_class, mentors.role, mentors.email, mentors.name, mentors.images, classes.class_name").
	Joins("left join classes on classes.id = mentors.id_class").
	Where("mentors.id = ?", id).Scan(&mentor).Error; err != nil {
		return admin.UserCore{}, err
	}
	cnv := ToDomainSingleMentor(mentor)
	return cnv, nil
	
}

func (ar *adminRepo) EditClass(input admin.ClassCore) (admin.ClassCore, error) {
	class := FromDomainUpdateClass(input)	

	// log.Print(input)
	if err := ar.db.Model(&class).Updates(&class).Error; err != nil {
		return admin.ClassCore{}, nil
	}
	
	return input, nil
}

