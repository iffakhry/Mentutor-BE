package delivery

import "be12/mentutor/features/mentor"

type UpdateUserResponse struct {
	ID uint
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Images string `json:"images" form:"images"`
}

func ToResponseUpdateUser(data mentor.UserCore) UpdateUserResponse{
	return UpdateUserResponse{
		ID: data.IdUser,
		Name: data.Name,
		Email: data.Email,
		Images: data.Images,
	}
}