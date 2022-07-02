package handle

import (
	"net/http"
	"time"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func GetTodoHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	today_start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	today_end := time.Date(year, month, day, 24, 0, 0, 0, time.Local)
	res, ok := model.SelecetTodoByUserIdAndTime(uint(id), today_start, today_end)
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func GetAllTodoHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	res, ok := model.SelecetTodoByUserId(uint(id))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": res})
	}
}

func InsertTodoHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	todo := model.Todo{}
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	todo.UserId = uint(id)
	todo.Time = time.Now().Local()
	ok := todo.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": todo})
	}
}

func DeleteTodoHandle(c *gin.Context) {
	todo := model.Todo{}
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	ok := todo.Delete()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": "OK"})
	}
}

func UpdateTodoHandle(c *gin.Context) {
	id := c.GetFloat64("id")
	todoJSON := model.Todo{}
	err := c.ShouldBindJSON(&todoJSON)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4001, util.ResMsg[4001])
		return
	}
	todo, ok := model.SelecetTodoByID(todoJSON.ID)
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
		return
	}
	todoJSON.Time = todo.Time
	todoJSON.UserId = uint(id)
	ok = todoJSON.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"data": todo})
	}
}
