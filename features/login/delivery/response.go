package delivery

import "be12/mentutor/features/login"

type LoginResponse struct {
	ID      uint   `json:"id_user"`
	Name    string `json:"name"`
	IdClass uint   `json:"id_class"`
	Class   string `json:"class"`
	Role    string `json:"role"`
	Images  string `json:"images"`
	Token   string `json:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {

	case "login":
		cnv := core.(login.Core)
		res = LoginResponse{ID: cnv.ID, Name: cnv.Name, IdClass: cnv.IdClass, Images: cnv.Images, Class: cnv.Class, Role: cnv.Role, Token: cnv.Token}
	}
	return res
}
