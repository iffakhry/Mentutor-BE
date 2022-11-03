package repository

import (
	"be12/mentutor/features/forum"

	"gorm.io/gorm"
)

type postData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) forum.DataInterface {
	return &postData{
		db: DB,
	}
}

func (pd *postData) Insert(data forum.Status, token int) (int, error) {

	dataModel := ToEntity(data)
	dataModel.IdMentee = uint(token)
	tx := pd.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (pd *postData) GetAll() ([]forum.Status, error) {
	var dataPost []Status
	tx := pd.db.Find(&dataPost)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// allComment := pd.commentDikit()
	dataPostUser := ToCoreArray(dataPost)

	return dataPostUser, nil
}

func (pd *postData) InsertStatus(data forum.Status, token int) (forum.Status, error) {

	dataModel := ToEntity(data)
	dataModel.IdMentee = uint(token)
	tx := pd.db.Create(&dataModel)
	if tx.Error != nil {
		return forum.Status{}, tx.Error
	}
	AddRes := dataModel.toPostUser()

	return AddRes, nil

}
