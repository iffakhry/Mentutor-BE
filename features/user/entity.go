package user

type Core struct {
	ID       uint
	Name     string
	Images   string
	Email    string
	Password string
	Role     string
	Class    string
}

type DataInterface interface {
	PostData(data Core) (row int, err error)
}

type UsecaseInterface interface {
	InsertData(data Core) (row int, err error)
}
