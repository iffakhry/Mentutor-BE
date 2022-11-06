package mentor

import "time"

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

type Class struct {
	ID        uint
	ClassName string
	Status    string
}

type Task struct {
	ID          uint
	IdClass     uint
	IdMentor    uint
	Description string
	File        string
	DueDate     time.Time
}

type Comment struct {
	ID       uint
	IdUser   uint
	IdStatus uint
	Caption  string 
}

type UsecaseInterface interface {
	UpdateProfile(input UserCore, role string) (UserCore, error)
}

type RepoInterface interface {
	EditProfileMentee(input UserCore) (UserCore, error)
	EditProfileMentor(input UserCore) (UserCore, error)
	GetSingleMentee(id uint) (UserCore, error)
	GetSingleMentor(id uint) (UserCore, error)
}
