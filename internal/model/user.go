package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Nama  string  `gorm:"column:nama;not null" json:"nama,omitempty"`
	Npm   string  `gorm:"column:npm;not null;unique" json:"npm,omitempty"`
	Tasks []Tasks `gorm:"foreignKey:UserId;references:ID" json:"tasks"`
}

func (Users) TableName() string {
	return "users"
}
