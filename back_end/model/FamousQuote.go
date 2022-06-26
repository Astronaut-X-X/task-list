package model

import "gorm.io/gorm"

type FamousQuote struct {
	UserId uint   `json:"user_id" gorm:"not null;"`
	Quote  string `json:"quote" gorm:"default '励志名言'"`
	Author uint   `json:"author" gorm:"default '作家'"`
	gorm.Model
}

func (item *FamousQuote) Create() (ok bool) {
	err := DB.Create(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *FamousQuote) Update() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Save(item).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (item *FamousQuote) Delete() (ok bool) {
	err := DB.Where("id = ?", item.Model.ID).Delete(&FamousQuote{}).Error
	if err != nil {
		return false
	}
	return true
}

func SelecetFamousQuoteByUserId(userid uint) (famousQuotes []FamousQuote, ok bool) {
	err := DB.Where("user_id = ?", userid).Find(&famousQuotes).Error
	if err != nil {
		// TODO log
		return famousQuotes, false
	}
	return famousQuotes, true
}
