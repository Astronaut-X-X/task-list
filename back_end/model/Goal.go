package model

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	UserId uint      `json:"userid" gorm:"not null;"`
	Cotent string    `json:"content" gorm:"default 'todu'"`
	Done   uint      `json:"done" gorm:"default 0"`
	Time   time.Time `json:"time" gorm:""`
	gorm.Model
}

func (goal *Goal) Create() (tx *gorm.DB, ok bool) {
	result := DB.Create(goal)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (goal *Goal) Update() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", goal.Model.ID).Save(goal)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (goal *Goal) Delete() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", goal.Model.ID).Delete(&Goal{})
	if result.Error == nil {
		return result, true
	}
	return result, false
}
