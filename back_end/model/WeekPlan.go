package model

import "gorm.io/gorm"

type WeekPlan struct {
	UserId      uint `json:"user_id" gorm:"not null;"`
	DailyPlanId uint `json:"daily_plan_id"`
	Week        uint `json:"week"`
	gorm.Model
}

func (item *WeekPlan) Create() (ok bool) {
	err := DB.Create(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *WeekPlan) Update() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Save(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *WeekPlan) Delete() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Delete(&WeekPlan{}).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func SelecetWeekPlanByUserId(userid uint) (weekPlans []WeekPlan, ok bool) {
	err := DB.Where("user_id = ?", userid).Find(&weekPlans).Error
	if err != nil {
		// TODO log
		return weekPlans, false
	}
	return weekPlans, true
}
