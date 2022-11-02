package class

import "time"

type CoreClass struct {
	ID         uint
	ClassName  string
	StatusEnum string
	UserID     uint
	TaskID     uint
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeletedAt  time.Time
}
