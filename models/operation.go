package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string
	Description string
	IsDone      bool `gorm:"default:false"`
	IsDeleted   bool `gorm:"default:false"`
	UserID      uint
}
