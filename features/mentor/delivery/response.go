package delivery

import "be12/mentutor/features/mentor"

type UpdateUserResponse struct {
	ID     uint
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Images string `json:"images" form:"images"`
}

type AddTaskResponse struct {
	ID          uint   `json:"id_task"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	File        string `json:"file"`
	DueDate     string `json:"due_date"`
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
