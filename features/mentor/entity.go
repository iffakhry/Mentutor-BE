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

type ClassCore struct {
	ID        uint
	ClassName string
	Status    string
}

type TaskCore struct {
	ID          uint
	IdClass     uint
	IdMentor    uint
	Title       string
	Description string
	File        string
	Images      string
	DueDate     time.Time
}

type CommentCore struct {
	ID       uint
	IdUser   uint
	IdStatus uint
	Caption  string
}

type UsecaseInterface interface {
	UpdateProfile(input UserCore, role string) (UserCore, error)
	AddTask(input TaskCore, role string) (TaskCore, error)
}

type RepoInterface interface {
	EditProfileMentee(input UserCore) (UserCore, error)
	EditProfileMentor(input UserCore) (UserCore, error)
	GetSingleMentee(id uint) (UserCore, error)
	GetSingleMentor(id uint) (UserCore, error)
	InsertTask(input TaskCore) (TaskCore, error)
}
