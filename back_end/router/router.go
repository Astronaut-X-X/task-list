package router

import (
	"github.com/Astronaut-X-X/TaskList/back_end/handler"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {

	R = gin.Default()
	R.Run()

	initRouter()

}

func initRouter() {
	api := R.Group("/api")
	v1 := api.Group("/v1")
	{
		user := v1.Group("/user", handler.Login)
		user.POST("/login")
		user.POST("") // 注册
	}

	{

	}

}

func initStaticFile() {
	R.Group("/")
}
