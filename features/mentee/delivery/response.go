package delivery

import (
	"be12/mentutor/features/mentee"
)

type UpdateResponse struct {
	IdUser uint   `json:"id_user"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Images string `json:"images"`
}

type StatusRespon struct {
	ID      uint   `json:"id"`
	Images  string `json:"images"`
	Caption string `json:"caption"`
	// Created_At time.Time `json:"created_at"`
	// Updated_At time.Time `json:"updated_at"`
}

func ToResponse(data mentee.Status) StatusRespon {
	return StatusRespon{
		ID:      data.ID,
		Caption: data.Caption,
		Images:  data.Images,
		// Created_At: data.CreatedAt,
	}
}

func ToCoreArray(pa []mentee.Status) []StatusRespon {
	var res []StatusRespon
	for _, val := range pa {
		res = append(res, StatusRespon{
			ID:      val.ID,
			Images:  val.Images,
			Caption: val.Caption,
		})
	}
	return res
}
func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
