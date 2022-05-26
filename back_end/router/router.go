package router

import (
	"github.com/Astronaut-X-X/TaskList/back_end/handler"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {

	R = gin.Default()
	initRouter()
	R.Run()

}

func initRouter() {
	api := R.Group("/api")
	v1 := api.Group("/v1")
	{
		user := v1.Group("/user", handler.LoginHandler())
		user.POST("/login")
		user.POST("") // 注册
	}
	v2 := api.Group("/v2")
	{
		v2.GET("/")
	}

}

func initStaticFile() {
	R.Group("/")
}
