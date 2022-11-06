package delivery

import "be12/mentutor/features/mentor"

type UpdateUserFormat struct {
	ID uint
	IdClass uint
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Images string `json:"images" form:"images"`
}

func ToDomainUpdateUser(data UpdateUserFormat) mentor.UserCore {
	return mentor.UserCore{
		IdUser: data.ID,
		IdClass: data.IdClass,
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		Images: data.Images,
	}
}