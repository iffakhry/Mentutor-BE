package login

type Core struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
	Class    string
	Token    string
}

type UsecaseInterface interface {
	Login(input Core) (Core, string, error)
}

type DataInterface interface {
	Login(input Core) (Core, error)
}
