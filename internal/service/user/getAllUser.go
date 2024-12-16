package user

import (
	"github.com/HaikalRFadhilahh/go-todos-list/internal/model"
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) ([]model.Users, error) {
	var d []model.Users
	err := db.Model(&model.Users{}).Find(&d).Error
	return d, err
}
