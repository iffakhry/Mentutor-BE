package delivery

import (
	"be12/mentutor/features/mentor"
)

type UpdateUserResponse struct {
	ID     uint
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Images string `json:"images" form:"images"`
}

type AddTaskResponse struct {
	ID          uint   `json:"id_task" form:"id_task"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images" form:"images"`
	File        string `json:"file" form:"file"`
	DueDate     string `json:"due_date" form:"due_date"`
}

type GetAllTask struct {
	IDTask      uint   `json:"id_task"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	FIle        string `json:"file"`
	DueDate     string `json:"due_date"`
}

type GetSingleTask struct {
	IdTask      uint               `json:"id_task"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Images      string             `json:"images"`
	File        string             `json:"file"`
	DueDate     string             `json:"due_date"`
	Submission  []SubmissionByTask `json:"submission"`
}

type SubmissionByTask struct {
	IdSubmission uint   `json:"id_submission"`
	Name         string `json:"name"`
	Score        int    `json:"score"`
	File         string `json:"file"`
}

func ToResponseUpdateUser(data mentor.UserCore) UpdateUserResponse {
	return UpdateUserResponse{
		ID:     data.IdUser,
		Name:   data.Name,
		Email:  data.Email,
		Images: data.Images,
	}
}

func ToResponseAddTask(data mentor.TaskCore) AddTaskResponse {
	return AddTaskResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Images:      data.Images,
		File:        data.File,
		DueDate:     data.DueDate.Format("2006-01-02 15:04 MST"),
	}
}

func ToResponseGetAllTask(data []mentor.TaskCore) []GetAllTask {
	var res []GetAllTask
	for _, val := range data {
		res = append(res, GetAllTask{
			IDTask:      val.ID,
			Title:       val.Title,
			Description: val.Description,
			Images:      val.Images,
			FIle:        val.File,
			DueDate:     val.DueDate.Format("2006-01-02 15:04 MST"),
		})
	}
	return res
}

func ToResponseSingleTask(task mentor.TaskCore, sub []mentor.SubmissionCore) GetSingleTask {
	var subs []SubmissionByTask

	for _, val := range sub {
		if val.IdTask == task.ID {
			subs = append(subs, SubmissionByTask{
				IdSubmission: val.ID,
				Name:         val.NameMentee,
				Score:        val.Score,
				File:         val.File,
			})
		}
	}

	return GetSingleTask{
		IdTask:      task.ID,
		Title:       task.Title,
		Description: task.Description,
		Images:      task.Images,
		File:        task.File,
		DueDate:     task.DueDate.Format("2006-01-02 15:04 MST"),
		Submission:  subs,
	}
}
