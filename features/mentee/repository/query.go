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
	if tx.Error != nil {
		return mentee.Status{}, tx.Error
	}
	log.Print(dataModel.Images, "INI DATAMODEL")
	AddRes := toPostUser(dataModel)
	AddRes.ID = dataModel.ID

	return AddRes, nil

}

func (md *menteeData) GetAllPosts() ([]mentee.Status, []mentee.CommentsCore, error) {
	var status []Status
	var comment []Comments

	tx := md.db.Model(&Status{}).Select("statuses.id, statuses.id_mentee, mentees.name, statuses.images, statuses.caption , mentees.role").
		Joins("left join mentees on mentees.id = statuses.id_mentee").Where("mentees.id = statuses.id_mentee").Scan(&status)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}
	cmn := md.db.Model(&Comments{}).Select("comments.id, comments.id_user, mentees.name,  comments.caption, comments.id_status , mentees.role").
		Joins("left join mentees on mentees.id = comments.id_user").Where("mentees.id = comments.id_user").Scan(&comment)
	if cmn.Error != nil {
		return nil, nil, cmn.Error
	}

	dataSC := toPostList(status)
	datacm := ToComent(comment)

	return dataSC, datacm, nil
}

func (md *menteeData) AddComment(data mentee.CommentsCore) (mentee.CommentsCore, error) {
	var input Comments
	input = ToEntityComent(data)
	res := md.db.Create(&input)
	if res.Error != nil {
		return mentee.CommentsCore{}, res.Error
	}
	// res = md.db.Model(&Comments{}).Select("name").Where("id = ?", input.IdStatus).Scan(&input)
	cnv := FromEntityComment(input)
	log.Print(input.Role, "    INI ROLE")
	return cnv, nil

}

func (md *menteeData) AddSub(data mentee.Submission) (mentee.Submission, error) {
	var input Submission
	input = FromEntitySub(data)

	res := md.db.Create(&input).Last(&input)
	log.Print(input.ID, " INI ID DARI QUERY")
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
