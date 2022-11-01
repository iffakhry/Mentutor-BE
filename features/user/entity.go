package user

type Core struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Images   string `json:"images" form:"images"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Class    string `json:"class" form:"class"`
}

type DataInterface interface {
	PostData(data Core) (row int, err error)
	MyProfile(token int) (data Core, err error)
	UpdateData(data Core) (row int, err error)
	GetAll() (data []Core, err error)
	SelectDataId(param, token int) (data Core, err error)
	DeleteData(token int) (int, error)
}

type UsecaseInterface interface {
	InsertData(data Core) (row int, err error)
	GetProfile(token int) (data Core, err error)
	PutDataId(data Core) (row int, err error)
	GetAlluser() (data []Core, err error)
	GetDataId(param, token int) (data Core, err error)
	Delete(token int) (int, error)
}
