package delivery

import "be12/mentutor/features/mentee"

type UpdateFormat struct {
	Name     string `form:"name" json:"name"`
	Email    string	`form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Images   string `form:"images"`
}

func ToEntity(data UpdateFormat) mentee.MenteeCore {
	return mentee.MenteeCore{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		Images: data.Password,
	}
}