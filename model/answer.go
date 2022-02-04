package model

import (
	"gorm.io/gorm"
)

type AnswerModel struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;auto_increment;not_null"`
	AnswerKey   string
	AnswerValue string
	IsDeleted   bool
	Event       EventModel `gorm:"foreignKey:EventId"`
	EventId     int
}

func (AnswerModel) TableName() string {
	return "data"
}
