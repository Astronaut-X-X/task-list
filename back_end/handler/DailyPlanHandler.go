package handler

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetDailyPlanHandler(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetDailyPlanByUserId(uint(id))
	if !ok {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	} else {
		util.Response(c, http.StatusOK, 5002, util.ResMsg[5002])
	}
}

func InsertDailyPlanHandler(c *gin.Context) {
	dailyPlan := model.DailyPlan{}
	err := c.ShouldBindJSON(&dailyPlan)
	if err != nil {
		util.Response(c, http.StatusOK, 4001, util.ResMsg[4001])
	}
	dailyPlan.UserId = uint(c.GetFloat64("id"))
	ok := dailyPlan.Create()
	if !ok {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyPlan})
	} else {
		util.Response(c, http.StatusOK, 5002, util.ResMsg[5002])
	}
}

func DeleteDailyPlanHandler(c *gin.Context) {
	dailyPlan := model.DailyPlan{}
	err := c.ShouldBindJSON(&dailyPlan)
	if err != nil {
		util.Response(c, http.StatusOK, 4001, util.ResMsg[4001])
	}
	ok := dailyPlan.Delete()
	if !ok {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	} else {
		util.Response(c, http.StatusOK, 5002, util.ResMsg[5002])
	}
}

func UpdateDailyPlanHandler(c *gin.Context) {
	dailyPlan := model.DailyPlan{}
	err := c.ShouldBindJSON(&dailyPlan)
	if err != nil {
		util.Response(c, http.StatusOK, 4001, util.ResMsg[4001])
	}
	ok := dailyPlan.Update()
	if !ok {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyPlan})
	} else {
		util.Response(c, http.StatusOK, 5002, util.ResMsg[5002])
	}
}
