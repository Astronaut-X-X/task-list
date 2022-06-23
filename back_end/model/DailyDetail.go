package model

import (
	"time"

	"gorm.io/gorm"
)

type DailyDetail struct {
	DailyPlanId uint      `json:"daily_plan_id" gorm:"not null;"`
	StartTime   time.Time `json:"start_time" gorm:"not null;"`
	EndTime     time.Time `json:"end_time" gorm:"not null;"`
	Detail      uint      `json:"detail" gorm:"not null;"`
	gorm.Model
}

func (item *DailyDetail) Create() (ok bool) {
	err := DB.Create(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *DailyDetail) Update() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Save(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *DailyDetail) Delete() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Delete(&DailyDetail{}).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}
