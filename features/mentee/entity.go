package mentee

type MenteeCore struct {
	IdUser uint
	Name string
	Email string
	Password string
	Images string
}

type UseCaseInterface interface {
	UpdateProfile(id uint, data MenteeCore) (MenteeCore, error)
}

type RepoInterface interface {
	EditProfile(id uint, data MenteeCore) (MenteeCore, error)
}