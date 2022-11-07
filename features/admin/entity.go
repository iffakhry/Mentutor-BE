package admin

type UserCore struct {
	IdUser   uint
	Name     string
	Email    string
	IdClass  uint
	Class    string
	Password string
	Role     string
	Images   string
}

type ClassCore struct {
	IdClass      uint
	ClassName    string
	Status       string
	TotalStudent int
}

type UsecaseInterface interface {
	AddUser(input UserCore, role string) (UserCore, error)
	GetAllUser(role string) ([]UserCore, []UserCore, error)
	AddNewClass(input ClassCore, role string) (ClassCore ,error)
	GetAllClass(role string) ([]ClassCore, error)
	UpdateUserAdmin(input UserCore, role string,) (UserCore, error)
	DeleteUser(id uint, role string) error
	GetSingleUser(id uint, role string) (UserCore, error)
	UpdateClass(input ClassCore, role string) (ClassCore, error)
	DeleteClass(id uint, role string) (error)
}

type RepoInterface interface {
	InsertMentee(input UserCore) (UserCore, error)
	InsertMentor(input UserCore) (UserCore, error)
	GetClass(id uint) (ClassCore, error)
	GetAllUser() ([]UserCore, []UserCore, error)
	InsertNewClass(input ClassCore) (ClassCore, error)
	GetAllClass() ([]ClassCore, error)
	EditUserMentee(input UserCore) (UserCore, error)
	EditUserMentor(input UserCore) (UserCore, error)
	DeleteUserMentee(id uint) error
	DeleteUserMentor(id uint) error
	GetSingleMentee(id uint) (UserCore, error)
	GetSingleMentor(id uint) (UserCore, error)
	EditClass(input ClassCore) (ClassCore, error)
	DeleteClass(id uint) (error)
}
