package main

import (
	"fmt"
	"github.com/Astronaut-X-X/GORM/controller"
	"github.com/Astronaut-X-X/GORM/model"
	"gorm.io/gorm"
)

func DoInsertUser() {
	user := model.User{
		Model:    gorm.Model{},
		Name:     "admin",
		Age:      1,
		UserName: "admin",
		Password: "123456",
	}

	if controller.InsertUser(&user) == 1 {
		fmt.Println("Insert user successful")
	}
}

func DoInsertMultipleUser() {
	for i := 1; i < 100; i++ {
		if controller.InsertUser(&model.User{
			Model:    gorm.Model{},
			Name:     fmt.Sprintf("User%v", i),
			Age:      uint8(i),
			UserName: fmt.Sprintf("User%v", i),
			Password: "123456",
		}) == 1 {
			fmt.Printf("Insert user %v successful \r\n", i)
		}
	}
}

func DoSelectUserById(id uint) {
	user := controller.SelectUserById(id)
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.UserName)
	fmt.Println(user.Password)
	fmt.Println(user.CreatedAt)
	fmt.Println(user.UpdatedAt)
	fmt.Println(user.DeletedAt)
}

func DoSelectUsers() {
	users := make([]model.User, 0)
	result := controller.SelectUserByUser(&users)
	fmt.Printf("Select users nums : %v \r\n", result)
	for _, v := range users {
		fmt.Println(v.ID)
		fmt.Println(v.Name)
		fmt.Println(v.Age)
		fmt.Println(v.UserName)
		fmt.Println(v.Password)
		fmt.Println(v.CreatedAt)
		fmt.Println(v.UpdatedAt)
		fmt.Println(v.DeletedAt)
	}
}

