package router

import (
	"fmt"
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/config"
	"github.com/Astronaut-X-X/TaskList/back_end/handler"
	m "github.com/Astronaut-X-X/TaskList/back_end/middleware"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	R = gin.Default()
	gin.SetMode(gin.DebugMode)
	R.MaxMultipartMemory = 8 << 20 // 8 MiB
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
	{
		R.GET("", handler.DefaultHandler)
		R.POST("", handler.DefaultHandler)
		R.PUT("", handler.DefaultHandler)
		R.DELETE("", handler.DefaultHandler)
		R.PATCH("", handler.DefaultHandler)
		R.HEAD("", handler.DefaultHandler)
		R.OPTIONS("", handler.DefaultHandler)
	}

	api := R.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/login", handler.LoginHandler)
		auth.POST("/register", handler.RegisterHandler)
		auth.POST("/flash", handler.FlashToken)
	}

	v1 := api.Group("/v1", m.VerifyToken())
	{
		user := v1.Group("/user")
		{
			user.GET("", handler.GetUserHandler)
			user.PUT("", handler.UpdateUserHandler)
			user.POST("/image", handler.UploadUserImageHandler)
		}

		dailyplan := v1.Group("/dailyplan")
		{
			dailyplan.GET("", handler.GetDailyPlanHandler)
			dailyplan.POST("", handler.InsertDailyPlanHandler)
			dailyplan.PUT("", handler.UpdateDailyPlanHandler)
			dailyplan.DELETE("", handler.DeleteDailyPlanHandler)
		}

		tasklist := v1.Group("/tasklsit")
		{
			tasklist.GET("", handler.DefaultHandler)
			tasklist.POST("")
			tasklist.PUT("")
			tasklist.DELETE("")
		}

		task := v1.Group("/task")
		{
			task.GET("")
			task.POST("")
			task.PUT("")
			task.DELETE("")
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

}

func initStaticFile() {
	R.Static(config.WEB_RELATIVE_PATH, config.WEB_STATIC_ROOT)
}
