package repository

import (
	"be12/mentutor/features/admin"
	"errors"
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
	
	if err := ar.db.Model(&Mentee{}).Select("classes.class_name, mentees.name, mentees.email, mentees.password, mentees.role, mentees.id").
	Where("classes.id = mentees.id_class").
	Joins("left join classes on classes.id = mentees.id_class").
	Scan(&mentees).Error; err != nil {
		return []admin.UserCore{}, []admin.UserCore{}, err
	}

	if err := ar.db.Model(&Mentor{}).Select("classes.class_name, mentors.name, mentors.email, mentors.password, mentors.role, mentors.id").
	Where("classes.id = mentors.id_class").
	Joins("left join classes on classes.id = mentors.id_class").
	Scan(&mentors).Error; err != nil {
		return []admin.UserCore{}, []admin.UserCore{}, err
	}

	cnvMentees := ToDomainMenteeArray(mentees)
	cnvMentors := ToDomainMentorArray(mentors)

	return cnvMentees, cnvMentors, nil
}

func (ar *adminRepo) InsertNewClass(input admin.ClassCore) (admin.ClassCore, error) {
	data := FromDomainClass(input)
	if err := ar.db.Create(&data).Last(&data).Error; err != nil {
		return admin.ClassCore{}, err
	}
	cnv := ToDomainClass(data)
	return cnv, nil
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
	var class Class
	var mentee Mentee
	
	data := FromDomainUpdateMentee(input)
	
	if err := ar.db.Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return admin.UserCore{}, err
	}

	ar.db.Where("id = ?", input.IdUser).Select("role, images").First(&mentee)
	ar.db.Where("id = ?", input.IdClass).Select("*"). First(&class)
	input.Class = class.ClassName
	input.Role = mentee.Role
	input.Images = mentee.Images
	return input, nil
}

func (ar *adminRepo) EditUserMentor(input admin.UserCore) (admin.UserCore, error) {
	var class Class
	var mentor Mentor

	data := FromDomainUpdateMentor(input)
	
	if err := ar.db.Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return admin.UserCore{}, err
	}
	ar.db.Where("id = ?", input.IdUser).Select("role, images").First(&mentor)
	ar.db.Where("id = ?", input.IdClass).Select("*"). First(&class)
	input.Class = class.ClassName
	input.Images = mentor.Images
	input.Role = mentor.Role
	return input, nil
}

func (ar *adminRepo) DeleteUserMentor(id uint) (error) {
	var mentor Mentor

	ar.db.Where("id_mentor = ?", id).Unscoped().Delete(&Task{})
	ar.db.Where("id_user = ?", id).Unscoped().Delete(&Comment{})

	if err := ar.db.Unscoped().Where("id = ?", id).Delete(&mentor).Error; err != nil {
		return err
	}

	return nil
}

func (ar *adminRepo) DeleteUserMentee(id uint) (error) {
	var mentee Mentee
	mentee.ID = id

	if err := ar.db.Unscoped().Delete(&mentee).Error; err != nil {
		return err
	}

	ar.db.Where("id_mentee = ?", id).Delete(&Submission{})
	ar.db.Where("id_user = ?", id).Delete(&Comment{})
	ar.db.Where("id_mentee = ?", id).Delete(&Status{})

	return nil
}

func (ar *adminRepo) GetSingleMentee(id uint) (admin.UserCore, error) {

	var mentee MenteeSingle

	
	if err := ar.db.Model(&Mentee{}).
	Select("mentees.id, mentees.id_class, mentees.role, mentees.email, mentees.name, mentees.images, classes.class_name").
	Joins("left join classes on classes.id = mentees.id_class").
	Where("mentees.id = ?", id).Scan(&mentee).Error; err != nil {
		mentee.ID = 0
		return admin.UserCore{}, err
	} 
	
	if mentee.Email == "" {
		
		mentee.ID = 0
	}
	cnv := ToDomainSingleMentee(mentee)

	return cnv, nil
} 

func (ar *adminRepo) GetSingleMentor(id uint) (admin.UserCore, error) {
	var mentor MentorSingle

	if err := ar.db.Model(&Mentor{}).
	Select("mentors.id, mentors.id_class, mentors.role, mentors.email, mentors.name, mentors.images, classes.class_name").
	Joins("left join classes on classes.id = mentors.id_class").
	Where("mentors.id = ?", id).Scan(&mentor).Error; err != nil {
		return admin.UserCore{}, err
	}

	if mentor.Email == "" {
		mentor.ID = 0
	}
	cnv := ToDomainSingleMentor(mentor)
	return cnv, nil	
}

func (ar *adminRepo) EditClass(input admin.ClassCore) (admin.ClassCore, error) {
	class := FromDomainUpdateClass(input)	

	if err := ar.db.Model(&class).Updates(&class).Error; err != nil {
		return admin.ClassCore{}, err
	}
	
	return input, nil
}

func (ar *adminRepo) DeleteClass(id uint) error {
	// var class Class
	log.Print(id)
	err := ar.db.Where("id = ?", id).Unscoped().Delete(&Class{})
	if  err.RowsAffected == 0 {
		return errors.New("class not found")
	}
	return nil
}

func (ar *adminRepo)GetSingleClass(id uint) (admin.ClassCore, error) {
	var class Class

	if err := ar.db.Where("id = ?", id).First(&class).Error; err != nil {
		return admin.ClassCore{} ,err
	}
	
	cnv := ToDomainClass(class)
	return cnv, nil
}