package model

import (
	"time"

	"gorm.io/gorm"
)

type Happy struct {
	UserId uint      `json:"user_id" gorm:"not null;"`
	Cotent string    `json:"content" gorm:"default 'Some thing make me happy'"`
	Time   time.Time `json:"time" gorm:""`
	gorm.Model
}

func (item *Happy) Create() (ok bool) {
	err := DB.Create(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *Happy) Update() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Save(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *Happy) Delete() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Delete(&Happy{}).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func SelecetHappyByUserIdAndTime(userid uint, start, end time.Time) (happys []Happy, ok bool) {
	err := DB.Where("user_id = ?", userid).Where("time BETWEEN ? AND ?", start, end).Find(&happys).Error
	if err != nil {
		// TODO log
		return happys, false
	}
	return happys, true
}

func SelecetHappyByID(id uint) (happy Happy, ok bool) {
	err := DB.Where("id = ?", id).First(&happy).Error
	if err != nil {
		// TODO log
		return happy, false
	}
	return happy, true
}

func SelecetHappyByUserId(userid uint) (happys []Happy, ok bool) {
	err := DB.Where("user_id = ?", userid).Order("time  desc").Find(&happys).Error
	if err != nil {
		// TODO log
		return happys, false
	}
	return happys, true
}
