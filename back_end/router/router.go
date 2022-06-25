package router

import (
	"fmt"
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/config"
	"github.com/Astronaut-X-X/TaskList/back_end/handle"
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
		R.GET("", handle.DefaultHandler)
		R.POST("", handle.DefaultHandler)
		R.PUT("", handle.DefaultHandler)
		R.DELETE("", handle.DefaultHandler)
		R.PATCH("", handle.DefaultHandler)
		R.HEAD("", handle.DefaultHandler)
		R.OPTIONS("", handle.DefaultHandler)
	}

	api := R.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/login", handle.LoginHandler)
		auth.POST("/register", handle.RegisterHandler)
		auth.POST("/flash", handle.FlashToken)
	}

	v1 := api.Group("/v1", m.VerifyToken())
	{
		user := v1.Group("/user")
		{
			user.GET("", handle.GetUserHandler)
			user.PUT("", handle.UpdateUserHandler)
			user.POST("/image", handle.UploadUserImageHandler)
		}

		dailyplan := v1.Group("/dailyplan")
		{
			dailyplan.GET("", handle.GetDailyPlanHandler)
			dailyplan.POST("", handle.InsertDailyPlanHandler)
			dailyplan.PUT("", handle.UpdateDailyPlanHandler)
			dailyplan.DELETE("", handle.DeleteDailyPlanHandler)
		}

		dailydetail := v1.Group("/dailydetail")
		{
			dailydetail.GET("", handle.GetDailyDetailHandler)
			dailydetail.POST("", handle.InsertDailyDetailHandler)
			dailydetail.PUT("", handle.UpdateDailyDetailHandler)
			dailydetail.DELETE("", handle.DeleteDailyDetailHandler)
		}

		weekplan := v1.Group("/weekplan")
		{
			weekplan.GET("", handle.GetWeekPlanHandler)
			weekplan.POST("", handle.InsertWeekPlanHandler)
			weekplan.PUT("", handle.UpdateWeekPlanHandler)
			weekplan.DELETE("", handle.DeleteWeekPlanHandler)
		}

		tasklist := v1.Group("/tasklsit")
		{
			tasklist.GET("", handle.DefaultHandler)
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
