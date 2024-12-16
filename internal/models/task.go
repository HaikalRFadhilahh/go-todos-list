package models

import (
	"time"

	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	Deadline time.Time `gorm:"column:deadline" json:"deadline,omitempty"`
	IsDone   int       `gorm:"column:isdone" json:"isDone,omitempty"`
	UserId   uint      `gorm:"column:userid" json:"userId,omitempty"`
	Users    Users     `gorm:"foreignKey:userid;references:id" json:"users,omitempty"`
}

func (Tasks) TableName() string {
	return "tasks"
}
