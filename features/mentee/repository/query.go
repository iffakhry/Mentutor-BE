package repository

import (
	"be12/mentutor/features/mentee"

	"github.com/labstack/gommon/log"
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

func (md *menteeData) AddStatus(data mentee.Status, token int) (mentee.Status, error) {

	dataModel := ToEntityMentee(data)
	dataModel.IdMentee = uint(token)
	tx := md.db.Create(&dataModel).Last(&dataModel)
	if tx.Error != nil {
		return mentee.Status{}, tx.Error
	}
	AddRes := toPostUser(dataModel)
	AddRes.ID = dataModel.ID

	return AddRes, nil

}

func (md *menteeData) GetAllPosts() ([]mentee.Status, []mentee.CommentsCore, []mentee.CommentsCore, error) {
	var status []Status
	var comment []Comments
	var mentorcom []Comments
	tx := md.db.Model(&Status{}).Select("statuses.id, statuses.id_mentee, mentees.name, statuses.images, statuses.caption , mentees.role").
		Joins("left join mentees on mentees.id = statuses.id_mentee").Where("mentees.id = statuses.id_mentee").Scan(&status)
	if tx.Error != nil {
		return nil, nil, nil, tx.Error
	}
	cmn := md.db.Model(&Comments{}).Select("comments.id, comments.id_user, mentees.name,mentees.role, comments.caption, comments.id_status ").
		Joins("left join mentees on mentees.id = comments.id_user").
		Where("mentees.id = comments.id_user ").Scan(&comment)
	if cmn.Error != nil {
		return nil, nil, nil, cmn.Error

	}
	com := md.db.Model(&Comments{}).Select("comments.id, comments.id_user, mentors.name,mentors.role, comments.caption, comments.id_status ").
		Joins("left join mentors on mentors.id = comments.id_user").
		Where("mentors.id = comments.id_user ").Scan(&mentorcom)
	if com.Error != nil {
		return nil, nil, nil, com.Error

	}

	dataSC := toPostList(status)
	datacm := ToComent(comment)
	comenmentor := ToComent(mentorcom)

	return dataSC, datacm, comenmentor, nil
}

func (md *menteeData) GetAllTask(idClass uint) ([]mentee.Task, error) {
	var task []Task

	tx := md.db.Model(&Task{}).Select("tasks.id,tasks.title, tasks.description, submissions.status, submissions.score, tasks.images, tasks.file, tasks.due_date").
		Joins("left join submissions on submissions.id_task = tasks.id").Where("tasks.id_class = ?", idClass).Scan(&task)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return toTaskList(task), nil
}

func (md *menteeData) AddComment(data mentee.CommentsCore) (mentee.CommentsCore, error) {
	var input Comments
	input = ToEntityComent(data)
	res := md.db.Create(&input)
	if res.Error != nil {
		return mentee.CommentsCore{}, res.Error
	}
	cnv := FromEntityComment(input)
	return cnv, nil

}

func (md *menteeData) AddSub(data mentee.Submission) (mentee.Submission, error) {
	var input Submission
	input = FromEntitySub(data)
	input.Status = "done"

	res := md.db.Create(&input).Last(&input)
	if res.Error != nil {
		log.Error("ERROR QUERY")
		return mentee.Submission{}, res.Error
	}
	res = md.db.Model(&Task{}).Select("title").Where("id = ?", input.IdTask).Scan(&input)
	cnv := ToEntitySub(input)

	return cnv, nil

}

func (md *menteeData) AddSubmis(param int, data mentee.Submission) (mentee.Submission, error) {
	// var input Submission
	input := FromEntitySub(data)

	res := md.db.Create(&input)
	if res.Error != nil {
		// log.Error("ERROR QUERY")
		return mentee.Submission{}, res.Error
	}
	log.Print(input.ID, " INI ID DARI QUERY")
	return data, nil
}

func (md *menteeData) GetSingleTask(idTask uint) (mentee.Task, error) {
	var res Task

	if err := md.db.Where("id = ?", idTask).First(&res).Error; err != nil {
		return mentee.Task{}, err
	}

	cnv := mentee.Task{ID: res.ID, DueDate: *res.DueDate}
	return cnv, nil
}

func (md *menteeData) InsertToken(data mentee.Token) (mentee.Token, error) {
	token := FromEntityToken(data)

	if err := md.db.Create(&token).Error; err != nil {
		return mentee.Token{}, err
	}
	return data, nil
}
