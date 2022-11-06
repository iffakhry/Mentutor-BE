package delivery

import (
	"be12/mentutor/features/mentee"
	"log"
)

type UpdateResponse struct {
	IdUser uint   `json:"id_user"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Images string `json:"images"`
}

type StatusRespon struct {
	ID       uint            `json:"id_status"`
	Images   string          `json:"images"`
	Caption  string          `json:"caption"`
	Comments []CommentRespon `json:"comments"`
}
type CommentRespon struct {
	ID      uint   `json:"id_comment"`
	Caption string `json:"caption"`
	Role    string `json:"role"`
	Name    string `json:"name"`
}

type SubResponse struct {
	ID    uint   `json:"id_submission"`
	Title string `json:"title"`
	File  string `json:"file"`
}

func ToResponse(data mentee.Status) StatusRespon {
	return StatusRespon{
		ID:      data.ID,
		Caption: data.Caption,
		Images:  data.Images,
	}
}

func ToResponseComments(data mentee.CommentsCore) CommentRespon {
	return CommentRespon{

		Caption: data.Caption,
	}
}

func ToResponseSub(data mentee.Submission) SubResponse {
	log.Print(data.ID, " INI ID RESPONSE")
	return SubResponse{

		ID:    data.ID,
		Title: data.Title,
		File:  data.File,
	}
}

func ToCoreArray(status []mentee.Status, coment []mentee.CommentsCore) []StatusRespon {
	var res []StatusRespon

	for i, val := range status {
		var comres []CommentRespon
		for j, v := range coment {
			if status[i].ID == coment[j].IdStatus {
				comres = append(comres, CommentRespon{ID: v.ID, Caption: v.Caption, Role: v.Role, Name: v.Name})

			}

		}

		res = append(res, StatusRespon{
			ID:       val.ID,
			Images:   val.Images,
			Caption:  val.Caption,
			Comments: comres,
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
