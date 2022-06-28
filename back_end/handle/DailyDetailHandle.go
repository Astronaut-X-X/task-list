package handle

import (
	"net/http"
	"strconv"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/service"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetDailyDetailHandle(c *gin.Context) {
	strId := c.Query("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 5004, util.ResMsg[5004])
		return
	}
	res, ok := model.SelecetDailyDetailByDailyPlanId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func GetTodayDailyDetailHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	dailyDetails, ok := service.GetTodayDailyDetailService(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyDetails})
	}
}

func InsertDailyDetailHandle(c *gin.Context) {
	dailyDetail := model.DailyDetail{}
	err := c.ShouldBindJSON(&dailyDetail)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	ok := dailyDetail.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyDetail})
	}
}

func DeleteDailyDetailHandle(c *gin.Context) {
	dailyDetail := model.DailyDetail{}
	err := c.ShouldBindJSON(&dailyDetail)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	ok := dailyDetail.Delete()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	}
}

func UpdateDailyDetailHandle(c *gin.Context) {
	dailyDetail := model.DailyDetail{}
	err := c.ShouldBindJSON(&dailyDetail)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	ok := dailyDetail.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": dailyDetail})
	}
}
