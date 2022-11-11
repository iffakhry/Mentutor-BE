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
	if err := mr.db.Model(&Submission{}).Select("submissions.id, submissions.file, submissions.id_task, mentees.name, submissions.score").
	Joins("left join mentees on mentees.id = submissions.id_mentee").
	Scan(&submission).Error; err != nil {
		return mentor.TaskCore{}, []mentor.SubmissionCore{}, err
	}
	
	cnvTask := ToDomainSingleTask(task)
	cnvSub := ToDomainTaskSub(submission)
	return cnvTask, cnvSub, nil
}

func(mr *mentorRepo) GetSingleTask(id uint) (mentor.TaskCore, error) {
	var res Task

	if err := mr.db.Where("id = ?", id).First(&res).Error; err != nil {
		return mentor.TaskCore{}, err 
	}

	cnv := ToDomainTask(res)
	return cnv, nil
}

func (mr *mentorRepo) EditTask(input mentor.TaskCore) (mentor.TaskCore, error) {
	data := FromDomainTask(input)

	if err := mr.db.Where("id = ?", data.ID).Updates(&data).Last(&data).Error; err != nil {
		return mentor.TaskCore{}, err
	}
	
	cnv := ToDomainTask(data)
	return cnv, nil
}

func (mr *mentorRepo) DeleteTask(idTask uint, FromidClass uint) (mentor.TaskCore, error) {

	err := mr.db.Where("id = ?", idTask).Delete(&Task{}) 
	if err.RowsAffected == 0 {
		return mentor.TaskCore{}, err.Error
	}
	return mentor.TaskCore{ID: idTask, IdClass: FromidClass}, nil
}

func (mr *mentorRepo) AddScore(input mentor.SubmissionCore) (mentor.SubmissionCore, error) {
	data := FromDomainSubmission(input)

	if err := mr.db.First(&data).Error; err != nil {
		return mentor.SubmissionCore{}, err
	}

	log.Print(data.Score)
	
	data.Score = input.Score
	if err := mr.db.Model(&Submission{}).Where("id = ?", input.ID).Updates(&data).Error; err != nil {
		return mentor.SubmissionCore{}, err
	}

	mr.db.Model(&Submission{}).Where("id = ?", input.ID).Select("file").Scan(&data)
	mr.db.Model(&Task{}).Where("id = ?", input.IdTask).Select("title").Scan(&data)
	cnv := ToDomainSubmission(data)
	return cnv, nil 
}

func (mr *mentorRepo) GetSubmission(id uint) error {

	if err := mr.db.Model(&Submission{}).Where("id = ?", id).Error; err != nil {
		log.Print("error")
		return err
	}
	return nil
}