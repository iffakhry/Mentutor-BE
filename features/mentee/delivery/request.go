package delivery

import "be12/mentutor/features/mentee"

type UpdateFormat struct {
	IdUser   uint
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Images   string `form:"images"`
}

type Request struct {
	ID      uint   `json:"id"`
	Caption string `json:"caption" form:"caption"`
	Images  string `json:"images" form:"images"`
}

func ToDomain(data Request) mentee.Status {
	return mentee.Status{
		// ID:      data.ID,
		Caption: data.Caption,
		Images:  data.Images,
	}
}

func ToEntity(data UpdateFormat) mentee.MenteeCore {
	return mentee.MenteeCore{
		IdUser:   data.IdUser,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Images:   data.Password,
	}
}
