package handle

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetWeekPlanHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetWeekPlanByUserId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func InsertWeekPlanHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	weekplan := model.WeekPlan{}
	err := c.ShouldBindJSON(&weekplan)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	weekplan.UserId = uint(id)
	ok := weekplan.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": weekplan})
	}
}

func DeleteWeekPlanHandle(c *gin.Context) {
	weekplan := model.WeekPlan{}
	err := c.ShouldBindJSON(&weekplan)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	ok := weekplan.Delete()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	}
}

func UpdateWeekPlanHandle(c *gin.Context) {
	weekplan := model.WeekPlan{}
	err := c.ShouldBindJSON(&weekplan)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	ok := weekplan.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": weekplan})
	}
}
