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
	log.Print(dataModel.ID, "ini id status")
	if tx.Error != nil {
		return mentee.Status{}, tx.Error
	}
	AddRes := toPostUser(dataModel)
	AddRes.ID = dataModel.ID

	return AddRes, nil

}

func (md *menteeData) GetAllPosts() ([]mentee.Status, []mentee.CommentsCore, error) {
	var status []Status
	var comment []Comments
	tx := md.db.Find(&status)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}
	cmn := md.db.Find(&comment)
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
	res = md.db.Model(&Mentee{}).Select("name").Where("id = ?", input.IdStatus).Scan(&input)
	cnv := FromEntityComment(input)
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
