package delivery

import "be12/mentutor/features/admin"

type RegisterFormat struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IdClass uint `json:"id_class" form:"id_class"`
	Role string `json:"role" form:"role"`
}


func ToDomain(data RegisterFormat) admin.UserCore {
	return admin.UserCore{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		IdClass: data.IdClass,
		Role: data.Role,
	}
}
