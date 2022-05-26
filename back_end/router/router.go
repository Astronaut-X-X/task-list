package router

import (
	"fmt"
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/config"
	"github.com/Astronaut-X-X/TaskList/back_end/handler"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	R = gin.Default()
	gin.SetMode(gin.DebugMode)
	initRouter()
	initStaticFile()
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", config.WEB_HOST, config.WEB_PORT),
		Handler:        R,
		ReadTimeout:    config.WEB_READTIMEOUT,
		WriteTimeout:   config.WEB_WRITETIMEOUT,
		MaxHeaderBytes: config.WEB_MAXHEADERBYTES,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic("Start web serve failed")
		}
	}()
}

func initRouter() {
	api := R.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", handler.LoginHandler)
		auth.POST("", handler.RegisterHandler)
	}

	v1 := api.Group("/v1")
	{

		task := v1.Group("/task")
		{
			task.GET("")
			task.POST("")
			task.PUT("")
			task.DELETE("")
		}

		tasklist := v1.Group("/tasklsit")
		{
			tasklist.GET("")
			tasklist.POST("")
			tasklist.PUT("")
			tasklist.DELETE("")
		}

		todo := v1.Group("/todo")
		{
			todo.GET("")
			todo.POST("")
			todo.PUT("")
			todo.DELETE("")
		}

		goal := v1.Group("/goal")
		{
			goal.GET("")
			goal.POST("")
			goal.PUT("")
			goal.DELETE("")
		}
	}

	// v2 TODO
	v2 := api.Group("/v2")
	{
		v2.GET("/")
	}

	// default_r := R.Group("/")
	// {
	// 	default_r.GET("", handler.HomeHandler)
	// 	default_r.POST("", handler.DefaultHandler)
	// 	default_r.PUT("", handler.DefaultHandler)
	// 	default_r.DELETE("", handler.DefaultHandler)
	// 	default_r.PATCH("", handler.DefaultHandler)
	// 	default_r.HEAD("", handler.DefaultHandler)
	// 	default_r.OPTIONS("", handler.DefaultHandler)
	// }

}

func initStaticFile() {
	R.Static(config.WEB_RELATIVE_PATH, config.WEB_STATIC_ROOT)
}
