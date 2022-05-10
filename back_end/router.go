package main

import (
	"github.com/Astronaut-X-X/GORM/api"
	"github.com/gin-gonic/gin"
)

//
//	InitRoute 创建 Gin 服务
//
//
func InitRoute(){
	r := gin.Default()

	defaultGroup := r.Group("/book")
	{
		defaultGroup.GET("/",api.IndexBooK)
	}

	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/home")
	}

	err := r.Run()
	if err != nil{
		panic("Web serve start failed.")
	}
}
