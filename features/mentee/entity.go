package mentee

import "time"

type MenteeCore struct {
	IdUser   uint
	Name     string
	Email    string
	Password string
	Images   string
}

type Status struct {
	ID        uint
	Caption   string
	Images    string
	IdMentee  uint
	CreatedAt time.Time
	// UpdateAt  time.Time
	// DeletedAt time.Time
	// Comment []CommentsCore
}

type UseCaseInterface interface {
	UpdateProfile(id uint, data MenteeCore) (MenteeCore, error)
	InsertStatus(data Status, token int) (Status, error)
	GetAll() (data []Status, err error)
}

type RepoInterface interface {
	EditProfile(id uint, data MenteeCore) (MenteeCore, error)
	AddStatus(data Status, token int) (Status, error)
	GetAllPosts() (data []Status, err error)
}
