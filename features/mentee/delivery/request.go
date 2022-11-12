package delivery

import (
	"be12/mentutor/features/mentee"

	"golang.org/x/oauth2"
)

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
type CommentFormat struct {
	IdStatus uint
	ID_User  uint
	Caption  string `json:"caption" form:"caption"`
}
type SubFormat struct {
	ID_Tasks  uint
	ID_Mentee uint
	File      string `json:"file" form:"file"`
}

func ToDomain(data Request) mentee.Status {
	return mentee.Status{
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

func ToDomainToken(code string, data *oauth2.Token) mentee.Token {
	return mentee.Token{
		Code:         code,
		AccessToken:  data.AccessToken,
		TokenType:    data.TokenType,
		RefreshToken: data.RefreshToken,
	}
}

// comments

func ToDomainComments(i CommentFormat) mentee.CommentsCore {
	return mentee.CommentsCore{
		IdStatus: i.IdStatus,
		ID_User:  i.ID_User,
		Caption:  i.Caption,
	}
}

func ToDomainSub(i SubFormat) mentee.Submission {
	return mentee.Submission{
		ID_Mentee: i.ID_Mentee,
		ID_Tasks:  i.ID_Tasks,
		File:      i.File,
	}
}
