package model

import "gorm.io/gorm"

type TaskList struct {
	UserId uint   `json:"userid" gorm:"not null;"`
	Cotent string `json:"content" gorm:"default 'todu'"`
	Color  string `json:"color" gorm:"default '#c2e9fb'"`
	gorm.Model
}

func (taskList *TaskList) Create() (tx *gorm.DB, ok bool) {
	result := DB.Create(taskList)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (taskList *TaskList) Update() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", taskList.Model.ID).Save(taskList)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (taskList *TaskList) Delete() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", taskList.Model.ID).Delete(&TaskList{})
	if result.Error == nil {
		return result, true
	}
	return result, false
}
