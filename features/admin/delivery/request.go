package delivery

type RegisterFormat struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IdClass uint `json:"id_class" form:"id_class"`
	Role string `json:"role" form:"role"`
}