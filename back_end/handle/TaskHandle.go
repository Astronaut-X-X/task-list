package handle

import (
	"net/http"
	"time"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetTaskHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	today_start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	today_end := time.Date(year, month, day, 24, 0, 0, 0, time.Local)
	res, ok := model.SelecetTaskByUserIdAndTime(uint(id), today_start, today_end)
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func GetAllTaskHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetTaskByUserId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func InsertTaskHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	task := model.Task{}
	err := c.ShouldBindJSON(&task)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	task.UserId = uint(id)
	task.Time = time.Now().Local()
	ok := task.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": task})
	}
}

func DeleteTaskHandle(c *gin.Context) {

	task := model.Task{}
	err := c.ShouldBindJSON(&task)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	ok := task.Delete()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	}
}

func UpdateTaskHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	taskJSON := model.Task{}
	err := c.ShouldBindJSON(&taskJSON)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
	}
	task, ok := model.SelecetTaskByID(taskJSON.ID)
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	}
	taskJSON.Time = task.Time
	taskJSON.UserId = uint(id)
	ok = taskJSON.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": task})
	}
}
