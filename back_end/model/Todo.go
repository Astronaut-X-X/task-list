package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	UserId uint      `json:"userid" gorm:"not null;"`
	Cotent string    `json:"content" gorm:"default 'todu'"`
	Done   uint      `json:"done" gorm:"default 0"`
	Time   time.Time `json:"time" gorm:""`
	gorm.Model
}

func (todo *Todo) Create() (tx *gorm.DB, ok bool) {
	result := DB.Create(todo)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (todo *Todo) Update() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", todo.Model.ID).Save(todo)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (todo *Todo) Delete() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", todo.Model.ID).Delete(&Todo{})
	if result.Error == nil {
		return result, true
	}
	return result, false
}
