package handle

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetWeekPlanHandler(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetWeekPlanByUserId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func InsertWeekPlanHandler(c *gin.Context) {
	dailyDetail := model.WeekPlan{}
	err := c.ShouldBindJSON(&dailyDetail)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	ok := dailyDetail.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyDetail})
	}
}

func DeleteWeekPlanHandler(c *gin.Context) {
	dailyDetail := model.WeekPlan{}
	err := c.ShouldBindJSON(&dailyDetail)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	ok := dailyDetail.Delete()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	}
}

func UpdateWeekPlanHandler(c *gin.Context) {
	dailyDetail := model.WeekPlan{}
	err := c.ShouldBindJSON(&dailyDetail)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	ok := dailyDetail.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyDetail})
	}
}
