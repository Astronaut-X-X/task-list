package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	UserId   uint      `json:"userid" gorm:"not null;"`
	Cotent   string    `json:"content" gorm:"default 'todu'"`
	Done     uint      `json:"done" gorm:"default 0"`
	Time     time.Time `json:"time" gorm:""`
	Meridiem string    `json:"meridiem" gorm:"default 'am'"`
	Color    string    `json:"color" gorm:"default '#c2e9fb'"`
	gorm.Model
}

func (task *Task) Create() (tx *gorm.DB, ok bool) {
	result := DB.Create(task)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (task *Task) Update() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", task.Model.ID).Save(task)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (task *Task) Delete() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", task.Model.ID).Delete(&Task{})
	if result.Error == nil {
		return result, true
	}
	return result, false
}
