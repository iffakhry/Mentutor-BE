package admin

type UserCore struct {
	IdUser   uint
	Name     string
	IdClass  uint
	Class    string
	Password string
	Role     string
}


type UsecaseInterface interface {
	AddUser(input UserCore) (UserCore, error)
}

type  RepoInterface interface {
	AddUser(input UserCore) (UserCore, error)
}