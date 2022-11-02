package migration

import (
	userModel "be12/mentutor/features/user/repository"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.Mentor{})

}
