package model

import "gorm.io/gorm"

type User struct {
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"unique;not null"`
	Model    gorm.Model
}

func (User) TableName() string {
	return "tl_user"
}

func CreateUser(user *User) (tx *gorm.DB, ok bool) {
	result := DB.Create(user)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func UpdateUser(id uint, user *User) (tx *gorm.DB, ok bool) {
	result := DB.Where("id = ?", id).Save(user)
	if result.RowsAffected > 0 {
		return result, true
	}
	return result, false
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
