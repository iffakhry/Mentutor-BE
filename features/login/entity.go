package login

type Core struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
	Images   string
	IdClass  uint
	Token    string
	Class    string
}

type UsecaseInterface interface {
	Login(input Core) (Core, string, error)
	GetToken(id uint) error
	InsertToken(idMentee uint) error
}

type DataInterface interface {
	Login(input Core) (Core, error)
	GetToken(id uint) error
	InsertToken(idMentee uint) error
}
