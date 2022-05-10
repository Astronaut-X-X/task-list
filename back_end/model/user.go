// model 包定义了系统的模型
package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      uint8  `json:"age";gorm:"default:18"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate user :", user.ID)
	return nil
}
