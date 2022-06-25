package handle

import (
	"net/http"
	"strconv"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetDailyDetailHandler(c *gin.Context) {
	strId := c.Query("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 5004, util.ResMsg[5004])
	}
	res, ok := model.SelecetDailyDetailByDailyPlanId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func InsertDailyDetailHandler(c *gin.Context) {
	dailyDetail := model.DailyDetail{}
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

func DeleteDailyDetailHandler(c *gin.Context) {
	dailyDetail := model.DailyDetail{}
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

func UpdateDailyDetailHandler(c *gin.Context) {
	dailyDetail := model.DailyDetail{}
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
