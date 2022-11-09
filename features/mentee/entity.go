package mentee

import "time"

type MenteeCore struct {
	IdUser   uint
	Name     string
	Email    string
	Password string
	Images   string
}

// FITUR STATUS
type Status struct {
	ID        uint
	Caption   string
	Name      string
	Images    string
	IdMentee  uint
	CreatedAt time.Time
	// UpdateAt  time.Time
	// DeletedAt time.Time
	Comments []CommentsCore //`json:",omitempty"`
}

// FITUR COMMENT
type CommentsCore struct {
	ID         uint
	ID_User    uint
	Name       string
	Role       string
	IdStatus   uint
	Caption    string
	Created_At time.Time
}

// FITUR SUBMISSION
type Submission struct {
	ID        uint
	ID_Mentee uint
	ID_Tasks  uint
	File      string
	Score     uint
	Title     string
	Status    string
}

type Task struct {
	ID          uint
	IdClass     uint
	IdMentor    uint
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	File        string    `gorm:"type:varchar(255);not null"`
	Images      string    `gorm:"type:varchar(255);not null"`
	DueDate     time.Time `gorm:"not null"`
	Score       uint
	Submissions []Submission `gorm:"foreignKey:IdTask"`
	Status      string
}
type UseCaseInterface interface {
	// UpdateProfile(id uint, data MenteeCore) (MenteeCore, error)
	InsertStatus(data Status, token int) (Status, error)
	GetAll() (data []Status, comen []CommentsCore, comenmentor []CommentsCore, err error)
	Insert(data CommentsCore) (CommentsCore, error)
	InsertSub(data Submission) (Submission, error)
	// InsertSubmis(param int, data Submission) (Submission, error)
	GetTask(idClass uint, role string) (data []Task, err error)
}

type RepoInterface interface {
	// EditProfile(id uint, data MenteeCore) (MenteeCore, error)
	AddStatus(data Status, token int) (Status, error)
	GetAllPosts() (data []Status, comen []CommentsCore, comenmentor []CommentsCore, err error)
	AddComment(data CommentsCore) (CommentsCore, error)
	AddSub(data Submission) (Submission, error)
	// AddSubmis(param int, data Submission) (Submission, error)
	GetAllTask(idClass uint) (data []Task, err error)
	GetSingleTask(idTask uint) (Task, error)
}
