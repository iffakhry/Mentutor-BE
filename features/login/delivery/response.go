package delivery

import "be12/mentutor/features/login"

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {

	case "login":
		cnv := core.(login.Core)
		res = LoginResponse{ID: cnv.ID, Name: cnv.Name, Class: cnv.Class, Role: cnv.Role, Token: cnv.Token}
	}
	return res
}
