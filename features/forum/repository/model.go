package repository

import (
	"be12/mentutor/features/forum"

	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	IdMentee uint   `json:"id_mentee" form:"Id_mentee"`
	Images   string `json:"images" form:"images"`
	Caption  string `json:"caption" form:"caption"`
}

func (p *Status) ToDomain() forum.Status {
	return forum.Status{
		ID:       p.ID,
		IdMentee: p.IdMentee,
		Caption:  p.Caption,
		Images:   p.Images,
	}
}

func ToEntity(data forum.Status) Status {
	return Status{
		IdMentee: data.IdMentee,
		Caption:  data.Caption,
		Images:   data.Images,
	}
}

func (dataPost *Status) toPostUser() forum.Status {

	dataPostCore := forum.Status{
		ID:       dataPost.ID,
		IdMentee: dataPost.IdMentee,
		Images:   dataPost.Images,
		Caption:  dataPost.Caption,
	}

	return dataPostCore

}

func ToCoreArray(pa []Status) []forum.Status {
	var res []forum.Status
	for _, val := range pa {
		res = append(res, forum.Status{
			ID:       val.ID,
			Images:   val.Images,
			Caption:  val.Caption,
			IdMentee: val.IdMentee,
		})
	}
	return res
}
