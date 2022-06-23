package model

import "gorm.io/gorm"

type DailyPlan struct {
	UserId uint `json:"user_id" gorm:"not null;"`
	Name   uint `json:"name" gorm:"not null;"`
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
	if err == nil {
		return false
	}
	return true
}

func SelecetDailyPlanByUserId(userid uint) (res []DailyDetail, ok bool) {
	err := DB.Where("id = ?", userid).Find(&res).Error
	if err != nil {
		// TODO log
		return res, false
	}
	return res, true
}
