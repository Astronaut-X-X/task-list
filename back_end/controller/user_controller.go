package controller

import (
	"fmt"
	. "github.com/Astronaut-X-X/GORM/global"
	"github.com/Astronaut-X-X/GORM/model"
)

func InsertUser(user *model.User) int {
	result := GlobalDB.Create(user)
	// GlobalDB.Select("name", "age","username","password").Create(*user)
	//GlobalDB.Omit("age").Create(*user)
	if result.Error != nil {
		// TODO record into log
		fmt.Println("InsertUser error :", result.Error)
	}
	return int(result.RowsAffected)
}

func SelectUserById(id uint) (user *model.User){
	GlobalDB.Where("id = ?", id).First(&user)
	return
}

func SelectUserByUser(users *[]model.User) int64{
	result := GlobalDB.Find(users)
	if result.Error != nil{
		fmt.Println("Select user by user error : ",result.Error)
	}
	return result.RowsAffected
}

func UpdateUserPasswordById(id int, password string) bool{
	result := GlobalDB.Model(&model.User{}).Where("id = ?",id).Update("password",password)
	if result.Error != nil{
		return false
	}else{
		return true
	}
}
