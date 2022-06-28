package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	UserId   uint      `json:"user_id" gorm:"not null;"`
	Cotent   string    `json:"content" gorm:"default 'todu'"`
	Done     bool      `json:"done" gorm:"default false"`
	Time     time.Time `json:"time" gorm:""`
	Meridiem string    `json:"meridiem" gorm:"default 'am'"`
	Color    string    `json:"color" gorm:"default '#c2e9fb'"`
	gorm.Model
}

func (item *Task) Create() (ok bool) {
	err := DB.Create(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *Task) Update() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Save(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *Task) Delete() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Delete(&Task{}).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func SelecetTaskByUserIdAndTime(userid uint, start, end time.Time) (tasks []Task, ok bool) {
	err := DB.Where("user_id = ?", userid).Where("time BETWEEN ? AND ?", start, end).Find(&tasks).Error
	if err != nil {
		// TODO log
		return tasks, false
	}
	return tasks, true
}

func SelecetTaskByID(id uint) (task Task, ok bool) {
	err := DB.Where("id = ?", id).First(&task).Error
	if err != nil {
		// TODO log
		return task, false
	}
	return task, true
}
