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
		R.GET("", handle.DefaultHandle)
		R.POST("", handle.DefaultHandle)
		R.PUT("", handle.DefaultHandle)
		R.DELETE("", handle.DefaultHandle)
		R.PATCH("", handle.DefaultHandle)
		R.HEAD("", handle.DefaultHandle)
		R.OPTIONS("", handle.DefaultHandle)
	}

	api := R.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/login", handle.LoginHandle)
		auth.POST("/register", handle.RegisterHandle)
		auth.POST("/flash", handle.FlashToken)
	}

	v1 := api.Group("/v1", m.VerifyToken())
	{
		user := v1.Group("/user")
		{
			user.GET("", handle.GetUserHandle)
			user.PUT("", handle.UpdateUserHandle)
			user.POST("/image", handle.UploadUserImageHandle)
		}

		dailyplan := v1.Group("/dailyplan")
		{
			dailyplan.GET("", handle.GetDailyPlanHandle)
			dailyplan.POST("", handle.InsertDailyPlanHandle)
			dailyplan.PUT("", handle.UpdateDailyPlanHandle)
			dailyplan.DELETE("", handle.DeleteDailyPlanHandle)
		}

		dailydetail := v1.Group("/dailydetail")
		{
			dailydetail.GET("", handle.GetDailyDetailHandle)
			dailydetail.GET("/today", handle.GetTodayDailyDetailHandle)
			dailydetail.POST("", handle.InsertDailyDetailHandle)
			dailydetail.PUT("", handle.UpdateDailyDetailHandle)
			dailydetail.DELETE("", handle.DeleteDailyDetailHandle)
		}

		weekplan := v1.Group("/weekplan")
		{
			weekplan.GET("", handle.GetWeekPlanHandle)
			weekplan.POST("", handle.InsertWeekPlanHandle)
			weekplan.PUT("", handle.UpdateWeekPlanHandle)
			weekplan.DELETE("", handle.DeleteWeekPlanHandle)
		}

		task := v1.Group("/task")
		{
			task.GET("", handle.GetAllTaskHandle)
			task.GET("/today", handle.GetTaskHandle)
			task.POST("", handle.InsertTaskHandle)
			task.PUT("", handle.UpdateTaskHandle)
			task.DELETE("", handle.DeleteTaskHandle)
		}

		tasklist := v1.Group("/tasklsit")
		{
			tasklist.GET("", handle.DefaultHandle)
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

}

func initStaticFile() {
	R.Static(config.WEB_RELATIVE_PATH, config.WEB_STATIC_ROOT)
}
