package delivery

import (
	"be12/mentutor/features/forum"
	"time"
)

type StatusRespon struct {
	ID         uint      `json:"id"`
	Images     string    `json:"images"`
	Caption    string    `json:"caption"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"update_at"`
}

func ToCoreArray(pa []StatusRespon) []forum.Status {
	var res []forum.Status
	for _, val := range pa {
		res = append(res, forum.Status{
			ID:      val.ID,
			Images:  val.Images,
			Caption: val.Caption,
		})
	}
	return res
}

func ToResponse(data forum.Status) StatusRespon {
	return StatusRespon{
		ID:      data.ID,
		Caption: data.Caption,
		Images:  data.Images,
	}
}
