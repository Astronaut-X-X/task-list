package model

import (
	"gorm.io/gorm"
)

type DailyDetail struct {
	DailyPlanId uint   `json:"daily_plan_id" gorm:"not null;"`
	StartTime   string `json:"start_time" gorm:"type:varchar(255);not null;"`
	EndTime     string `json:"end_time" gorm:"type:varchar(255);not null;"`
	Detail      string `json:"detail" gorm:"type:varchar(255);not null;"`
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

func SelecetDailyDetailByDailyPlanId(dailyPlanid uint) (dailyDetails []DailyDetail, ok bool) {
	err := DB.Where("daily_plan_id = ?", dailyPlanid).Find(&dailyDetails).Error
	if err != nil {
		// TODO log
		return dailyDetails, false
	}
	return dailyDetails, true
}