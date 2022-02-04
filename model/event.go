package model

import (
	"gorm.io/gorm"
)

type EventModel struct {
	gorm.Model
	ID   uint `gorm:"primaryKey;auto_increment;not_null"`
	Name string
}

func (EventModel) TableName() string {
	return "events"
}
