package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	UserId uint      `json:"user_id" gorm:"not null;"`
	Cotent string    `json:"content" gorm:"default 'todu'"`
	Done   bool      `json:"done" gorm:"default 0"`
	Time   time.Time `json:"time" gorm:""`
	gorm.Model
}

func (item *Todo) Create() (ok bool) {
	err := DB.Create(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *Todo) Update() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Save(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *Todo) Delete() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Delete(&Todo{}).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func SelecetTodoByUserIdAndTime(userid uint, start, end time.Time) (todos []Todo, ok bool) {
	err := DB.Where("user_id = ?", userid).Where("time BETWEEN ? AND ?", start, end).Find(&todos).Error
	if err != nil {
		// TODO log
		return todos, false
	}
	return todos, true
}

func SelecetTodoByID(id uint) (todo Todo, ok bool) {
	err := DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		// TODO log
		return todo, false
	}
	return todo, true
}

func SelecetTodoByUserId(userid uint) (todos []Todo, ok bool) {
	err := DB.Where("user_id = ?", userid).Order("time  desc").Find(&todos).Error
	if err != nil {
		// TODO log
		return todos, false
	}
	return todos, true
}
