package main

import (
	"github.com/Astronaut-X-X/GORM/config"
	"github.com/Astronaut-X-X/GORM/global"
	"github.com/Astronaut-X-X/GORM/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func InitDB() {
	var err error

	if global.GlobalDB, err = gorm.Open(sqlite.Open(config.DB_NAME),&gorm.Config{});err != nil{
		panic("Connect sqlite failed.")
	}

	if global.SqlDB,err = global.GlobalDB.DB();err != nil{
		panic("Connect sqlite failed.")
	}

	global.SqlDB.SetMaxOpenConns(100)          // 设置做多连接
	global.SqlDB.SetMaxIdleConns(10)           // 设置闲置时最少连接
	global.SqlDB.SetConnMaxIdleTime(time.Hour) // 设置闲置最长存活时间
	global.SqlDB.SetConnMaxLifetime(time.Hour) // 设置最长存活时间
}

func InitAutoMigrate(){
	err := global.GlobalDB.AutoMigrate(&model.User{},&model.Admin{},&model.Book{},&model.BookInfo{})
	if err != nil{
		panic("AutoMigrate failed.")
	}
	model.InitAdminRole()
}