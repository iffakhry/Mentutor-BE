package repository

import (
	"be12/mentutor/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	Images   string `gorm:"type:varchar(255);not null"`
	IdClass  uint
}

func FromEntity(data mentee.MenteeCore) Mentee {
	return Mentee{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Images:   data.Images,
	}
}

func ToEntity(id uint, data Mentee) mentee.MenteeCore {
	return mentee.MenteeCore{
		IdUser: id,
		Name:   data.Name,
		Email:  data.Email,
		Images: data.Images,
	}
}

type Status struct {
	gorm.Model
	IdMentee uint   `json:"id_mentee" form:"Id_mentee"`
	Images   string `json:"images" form:"images"`
	Caption  string `json:"caption" form:"caption"`
}

func (p *Status) ToDomain() mentee.Status {
	return mentee.Status{
		ID:       p.ID,
		IdMentee: p.IdMentee,
		Caption:  p.Caption,
		Images:   p.Images,
	}
}

func ToEntityMentee(data mentee.Status) Status {
	return Status{
		Model:    gorm.Model{ID: data.ID},
		IdMentee: data.IdMentee,
		Caption:  data.Caption,
		Images:   data.Images,
	}
}

func toPostUser(dataPost Status) mentee.Status {

	dataPostCore := mentee.Status{
		ID:       dataPost.ID,
		IdMentee: dataPost.IdMentee,
		Images:   dataPost.Images,
		Caption:  dataPost.Caption,
	}

	return dataPostCore

}

func ToCoreArray(pa []Status) []mentee.Status {
	var res []mentee.Status
	for _, val := range pa {
		res = append(res, mentee.Status{
			ID:        val.ID,
			Images:    val.Images,
			Caption:   val.Caption,
			IdMentee:  val.IdMentee,
			CreatedAt: val.CreatedAt,
		})
	}
	return res
}
