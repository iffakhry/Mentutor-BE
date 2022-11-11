package delivery

import "be12/mentutor/features/login"

type LoginResponse struct {
	ID        uint   `json:"id_user"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IdClass   uint   `json:"id_class"`
	Class     string `json:"class_name"`
	Role      string `json:"role"`
	Images    string `json:"images"`
	Token     string `json:"token"`
	AuthGmail string `json:"auth_gmail"`
}

func ToResponse(core interface{}, code string, url string) interface{} {
	var res interface{}
	switch code {

	case "login":
		cnv := core.(login.Core)
		res = LoginResponse{ID: cnv.ID, Name: cnv.Name, Email: cnv.Email, IdClass: cnv.IdClass, Images: cnv.Images, Class: cnv.Class, Role: cnv.Role, Token: cnv.Token, AuthGmail: url}
	}
	return res
}
