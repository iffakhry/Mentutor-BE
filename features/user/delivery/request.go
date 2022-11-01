package delivery

import "be12/mentutor/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Class    string `json:"class" form:"class"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Images   string `json:"images" form:"images"`
	Role     string `json:"role" form:"role"`
	Saldo    uint   `json:"saldo" form:"saldo"`
}

// type UpdateFormat struct {
// 	ID       uint
// 	Name     string `form:"name" json:"name"`
// 	Saldo    int    `form:"saldo" json:"saldo"`
// 	Password string `form:"password" json:"password"`
// 	Email    string `form:"email" json:"email"`
// 	Images   string `form:"images" json:"images"`
// 	Phone    int    `form:"phone" json:"phone"`
// 	Bio      string `form:"bio" json:"bio"`
// }

func toCore(data UserRequest) user.Core {
	return user.Core{
		Name:     data.Name,
		Images:   data.Images,
		Email:    data.Email,
		Class:    data.Class,
		Password: data.Password,
		Role:     data.Role,
	}
}
