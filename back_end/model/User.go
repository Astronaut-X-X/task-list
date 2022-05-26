package model

import (
	"github.com/Astronaut-X-X/TaskList/back_end/util"

	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"unique;not null;"`
	Password string `json:"password" gorm:"not null;"`
	Image    string `json:"image" gorm:"default ''"`
	Email    string `json:"email" gorm:"not null;"`
	gorm.Model
}

func (User) TableName() string {
	return "tl_user"
}

func (user *User) Create() (tx *gorm.DB, ok bool) {
	result := DB.Create(user)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (user *User) Update() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", user.Model.ID).Save(user)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (user *User) Delete() (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", user.Model.ID).Delete(&User{})
	if result.Error == nil {
		return result, true
	}
	return result, false
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = util.MD5Encipher(user.Password)
	// TODO logging
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.Password = util.MD5Encipher(user.Password)
	// TODO logging
	return nil
}

func SelectUserByUsername(username string) (user User, ok bool) {
	result := DB.Model(&User{}).Where("username = ?", username).First(&user)
	if result.RowsAffected > 0 {
		return user, true
	}
	return user, false
}

func SelectUserById(id uint) (user User, ok bool) {
	result := DB.First(&user, id)
	if result.RowsAffected > 0 {
		return user, true
	}
	return user, false
}

func SelectUser(id uint) (users []User, ok bool) {
	result := DB.Find(&users, id)
	if result.RowsAffected > 0 {
		return users, true
	}
	return users, false
}
