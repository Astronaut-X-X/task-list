package model

import (
	"gorm.io/gorm"
)

type DailyPlan struct {
	UserId uint   `json:"user_id" gorm:"not null;"`
	Name   string `json:"name" gorm:"type:varchar(255);not null;"`
	gorm.Model
}

func (item *DailyPlan) Create() (ok bool) {
	err := DB.Create(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *DailyPlan) Update() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Save(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *DailyPlan) Delete() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Delete(&DailyPlan{}).Error
	if err != nil {
		return false
	}
	return true
}

func SelecetDailyPlanByUserId(userid uint) (dailyPlans []DailyPlan, ok bool) {
	err := DB.Where("user_id = ?", userid).Find(&dailyPlans).Error
	if err != nil {
		// TODO log
		return dailyPlans, false
	}
	return dailyPlans, true
}
