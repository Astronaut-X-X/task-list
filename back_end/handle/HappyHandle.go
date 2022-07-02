package handle

import (
	"net/http"
	"time"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetHappyHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	today_start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	today_end := time.Date(year, month, day, 24, 0, 0, 0, time.Local)
	res, ok := model.SelecetHappyByUserIdAndTime(uint(id), today_start, today_end)
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func GetAllHappyHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetHappyByUserId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func GetHappyPageHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetHappyByUserId(uint(id))
	happy := model.Happy{}
	err := c.ShouldBindJSON(&happy)
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func InsertHappyHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	happy := model.Happy{}
	err := c.ShouldBindJSON(&happy)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	happy.UserId = uint(id)
	happy.Time = time.Now().Local()
	ok := happy.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": happy})
	}
}

func DeleteHappyHandle(c *gin.Context) {
	happy := model.Happy{}
	err := c.ShouldBindJSON(&happy)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	ok := happy.Delete()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	}
}

func UpdateHappyHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	happyJSON := model.Happy{}
	err := c.ShouldBindJSON(&happyJSON)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	happy, ok := model.SelecetHappyByID(happyJSON.ID)
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
		return
	}
	happyJSON.Time = happy.Time
	happyJSON.UserId = uint(id)
	ok = happyJSON.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": happy})
	}
}
