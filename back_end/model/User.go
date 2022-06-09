package model

import (
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"unique;not null;"`
	Password string `json:"password" gorm:"not null;"`
	Image    string `json:"image" gorm:"default ''"`
	Email    string `json:"email" gorm:"not null;"`
	gorm.Model
}

func (user *User) Create() (ok bool) {
	err := DB.Create(user).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (user *User) Update() (ok bool) {
	err := DB.Where("id = ?", user.Model.ID).Save(user).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func (user *User) Delete() (ok bool) {
	err := DB.Where("id = ?", user.Model.ID).Delete(&User{}).Error
	if err != nil {
		// TODO log
		return false
	}
	return true
}

func SelectUserByUsername(username string) (user User, ok bool) {
	err := DB.Model(&User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		// TODO log
		return user, false
	}
	return user, true
}

func SelectUserById(id int) (user User, ok bool) {
	err := DB.First(&user, id).Error
	if err != nil {
		// TODO log
		return user, false
	}
	return user, true
}

func SelectUser() (users []User, ok bool) {
	err := DB.Find(&users).Error
	if err != nil {
		return users, false
	}
	return users, true
}
