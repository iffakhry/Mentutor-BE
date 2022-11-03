package forum

import "time"

type Status struct {
	ID       uint
	Caption  string
	Images   string
	IdMentee uint
	// CreatedAt time.Time
	// UpdateAt  time.Time
	// DeletedAt time.Time
	Comment []CommentsCore
}

type Mentee struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Hp       string `json:"hp" form:"hp"`
	Bio      string `json:"bio" form:"bio"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Status   []Status
}
type CommentsCore struct {
	ID         uint
	User_ID    uint
	Post_ID    uint
	Comment    string `json:"comment" form:"comment"`
	Created_At time.Time
}
type DataInterface interface {
	Insert(data Status, token int) (int, error)
	GetAll() (data []Status, err error)
	InsertStatus(data Status, token int) (Status, error)
}

type UseCaseInterface interface {
	AddPost(data Status, token int) (int, error)
	GetAllPosts() (data []Status, err error)
	AddStatus(data Status, token int) (Status, error)
}
