package delivery

import "be12/mentutor/features/admin"

type RegisterResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	IdClass uint `json:"id_class"`
	Role string `json:"role"`
}

func ToResponse(data admin.UserCore) RegisterResponse {
	return RegisterResponse{
		Name: data.Name,
		Email: data.Email,
		IdClass: data.IdClass,
		Role: data.Role,
	}
}