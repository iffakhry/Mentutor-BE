package repository

import (
	"be12/mentutor/features/mentor"
	"log"

	"gorm.io/gorm"
)

type mentorRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentor.RepoInterface {
	return &mentorRepo{
		db: db,
	}
}

func (mr *mentorRepo) GetSingleMentee(id uint) (mentor.UserCore, error) {

	var mentee Mentee
	mentee.ID = id

	if err := mr.db.Model(&Mentee{}).
	Select("mentees.id, mentees.id_class, mentees.role, mentees.email, mentees.name, mentees.images, classes.class_name").
	Joins("left join classes on classes.id = mentees.id_class").
	Where("mentees.id = ?", id).Scan(&mentee).Error; err != nil {
		return mentor.UserCore{}, err
	}
	cnv := ToDomainMentee(mentee)
	return cnv, nil
} 

func (mr *mentorRepo) GetSingleMentor(id uint) (mentor.UserCore, error) {
	var res Mentor
	res.ID = id

	if err := mr.db.Model(&Mentor{}).
	Select("mentors.id, mentors.id_class, mentors.role, mentors.email, mentors.name, mentors.images, classes.class_name").
	Joins("left join classes on classes.id = mentors.id_class").
	Where("mentors.id = ?", id).Scan(&res).Error; err != nil {
		return mentor.UserCore{}, err
	}
	cnv := ToDomainMentor(res)
	return cnv, nil	
}

func (mr *mentorRepo) EditProfileMentee(input mentor.UserCore) (mentor.UserCore, error) {
	data := FromDomainMentee(input)

	if err := mr.db.Model(&Mentee{}).Where("id = ?", input.IdUser).Updates(&data).Error; err != nil {
		return mentor.UserCore{}, err
	}
	return input, nil
}

func (mr *mentorRepo) EditProfileMentor(input mentor.UserCore) (mentor.UserCore, error) {
	data := FromDomainMentor(input)

	if err := mr.db.Model(&Mentor{}).Where("id = ?", input.IdUser).Updates(&data).Error; err != nil {
		return mentor.UserCore{}, err
	}
	return input, nil
}

func (mr *mentorRepo) InsertTask(input mentor.TaskCore) (mentor.TaskCore, error) {
	data := FromDomainTask(input)

	if err := mr.db.Create(&data).Last(&data).Error; err != nil {
		return mentor.TaskCore{}, err
	}

	cnv := ToDomainTask(data)
	return cnv, nil
}

func (mr *mentorRepo) GetAllTask() ([]mentor.TaskCore, error) {
	var tasks []Task

	if err := mr.db.Find(&tasks).Error; err != nil {
		return []mentor.TaskCore{}, err
	}

	log.Print(tasks)

	cnv := ToDomainAllTask(tasks)
	return cnv, nil
}

func (mr *mentorRepo) GetTaskSub(idTask uint) (mentor.TaskCore, []mentor.SubmissionCore, error) {
	var task Task
	var submission []Submission

	if err := mr.db.Where("id = ?", idTask).First(&task).Error; err != nil {
		return mentor.TaskCore{}, []mentor.SubmissionCore{}, err
	}
	if err := mr.db.Find(&submission).Joins("left join mentees on mentees.id = submissions.id_mentee").Scan(&submission).Error; err != nil {
		return mentor.TaskCore{}, []mentor.SubmissionCore{}, err
	}
	
	cnvTask := ToDomainSingleTask(task)
	cnvSub := ToDomainTaskSub(submission)
	return cnvTask, cnvSub, nil
}