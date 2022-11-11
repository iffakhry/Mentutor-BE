package mentee

import (
	"time"

)

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
	Token     string
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

type Token struct {
	Code         string
	AccessToken  string    
	TokenType    string    
	RefreshToken string   
	Expiry       time.Time
}

type UseCaseInterface interface {
	InsertStatus(data Status, token int) (Status, error)
	GetAll() (data []Status, comen []CommentsCore, comenmentor []CommentsCore, err error)
	Insert(data CommentsCore) (CommentsCore, error)
	InsertSub(data Submission) (Submission, error)
	GetTask(idClass uint, role string) (data []Task, err error)
	AddToken(token Token) (Token, error)
}

type RepoInterface interface {
	AddStatus(data Status, token int) (Status, error)
	GetAllPosts() (data []Status, comen []CommentsCore, comenmentor []CommentsCore, err error)
	AddComment(data CommentsCore) (CommentsCore, error)
	AddSub(data Submission) (Submission, error)
	GetAllTask(idClass uint) (data []Task, err error)
	GetSingleTask(idTask uint) (Task, error)
	InsertToken(token Token) (Token, error)
}
