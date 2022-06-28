package handle

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetDailyPlanHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetDailyPlanByUserId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func InsertDailyPlanHandle(c *gin.Context) {
	dailyPlan := model.DailyPlan{}
	err := c.ShouldBindJSON(&dailyPlan)
	if err != nil {
		util.Response(c, http.StatusOK, 4001, util.ResMsg[4001])
	}
	dailyPlan.UserId = uint(c.GetFloat64("id"))
	ok := dailyPlan.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyPlan})
	}
}

func DeleteDailyPlanHandle(c *gin.Context) {
	dailyPlan := model.DailyPlan{}
	err := c.ShouldBindJSON(&dailyPlan)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	ok := dailyPlan.Delete()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	}
}

func UpdateDailyPlanHandle(c *gin.Context) {
	dailyPlan := model.DailyPlan{}
	err := c.ShouldBindJSON(&dailyPlan)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	dailyPlan.UserId = uint(c.GetFloat64("id"))
	ok := dailyPlan.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyPlan})
	}
}
