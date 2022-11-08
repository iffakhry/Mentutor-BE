package migration

import (
	// userModel "be12/mentutor/features/user/repository"

	"time"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"type:varchar(255);not null;unique"`
	Password    string `gorm:"type:varchar(255);not null"`
	Images      string `gorm:"type:varchar(255)"`
	Role        string `gorm:"type:enum('mentee');not null"`
	IdClass     uint
	Submissions []Submission `gorm:"foreignKey:IdMentee"`
	Statuses    []Status     `gorm:"foreignKey:IdMentee"`
}

type Mentor struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Images   string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:enum('admin','mentor');not null"`
	IdClass  uint
	Tasks    []Task    `gorm:"foreignKey:IdMentor"`
}

type Class struct {
	gorm.Model
	ClassName string   `gorm:"type:varchar(255);unique;not null"`
	Status    string   `gorm:"type:enum('active','non_active');not null"`
	Mentors   []Mentor `gorm:"foreignKey:IdClass"`
	Mentees   []Mentee `gorm:"foreignKey:IdClass"`
	Tasks     []Task   `gorm:"foreignKey:IdClass"`
}

type Task struct {
	gorm.Model
	IdClass     uint
	IdMentor    uint
	Title       string       `gorm:"type:varchar(255);not null"`
	Description string       `gorm:"type:varchar(255);not null"`
	File        string       `gorm:"type:varchar(255);not null"`
	Images      string       `gorm:"type:varchar(255);not null"`
	DueDate     *time.Time   `gorm:"not null"`
	Submissions []Submission `gorm:"foreignKey:IdTask"`
}

type Submission struct {
	gorm.Model
	IdMentee uint
	IdTask   uint
	// Title    string `gorm:"type:varchar(255);not null"`
	File  string `gorm:"type:varchar(255);not null"`
	Score int    `gorm:"type:int(3);not null"`
}

type Status struct {
	gorm.Model
	IdMentee uint
	Caption  string    `gorm:"type:varchar(255);not null"`
	Images   string    `gorm:"type:varchar(255);not null"`
	Comment  []Comment `gorm:"foreignKey:IdStatus"`
}

type Comment struct {
	gorm.Model
	IdUser   uint
	IdStatus uint
	Caption  string `gorm:"type:varchar(255);not null"`
}

func InitMigrate(db *gorm.DB) {
	// db.AutoMigrate(&userModel.Mentor{})
	db.AutoMigrate((&Mentor{}))
	db.AutoMigrate((&Mentee{}))
	db.AutoMigrate((&Class{}))
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Submission{})
	db.AutoMigrate(&Status{})
	db.AutoMigrate(&Comment{})
}
