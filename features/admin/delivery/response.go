package delivery

import "be12/mentutor/features/admin"

type RegisterResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	IdClass uint `json:"id_class"`
	Role string `json:"role"`
}

type GetUserResponse struct {
	IdUser uint `json:"id_user"`
	Name string `json:"name"`
	Role string `json:"role"`
	Class string `json:"class"`
}

type GetAllClassResponse struct {
	IdClass uint `json:"id_class"`
	Class string `json:"class_name"`
	Status string `json:"status"`
	TotalStudent int `json:"total_student"`
}

func ToResponse(data admin.UserCore) RegisterResponse {
	return RegisterResponse{
		Name: data.Name,
		Email: data.Email,
		IdClass: data.IdClass,
		Role: data.Role,
	}
}

func ToResponseUserArray(mentee []admin.UserCore, mentor []admin.UserCore) []GetUserResponse {
	var res []GetUserResponse
	for _, val := range mentee{
		res = append(res, GetUserResponse{
			IdUser: val.IdUser,
			Name: val.Name,
			Role: val.Role,
			Class: val.Class,
		})
	}
	for _, val := range mentor{
		res = append(res, GetUserResponse{
			IdUser: val.IdUser,
			Name: val.Name,
			Role: val.Role,
			Class: val.Class,
		})
	}
	return res
}

func ToResponseClassArray(data []admin.ClassCore) []GetAllClassResponse {
	var res []GetAllClassResponse

	for _, val := range data {
		res = append(res, GetAllClassResponse{
			IdClass: val.IdClass,
			Class: val.ClassName,
			Status: val.Status,
			TotalStudent: val.TotalStudent,
		})
	}
	return res
}