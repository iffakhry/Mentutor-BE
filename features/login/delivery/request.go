package delivery

import "be12/mentutor/features/login"

type LoginReq struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) login.Core {
	switch i.(type) {

	case LoginFormat:
		cnv := i.(LoginFormat)
		return login.Core{Email: cnv.Email, Password: cnv.Password}
	}
	return login.Core{}
}
