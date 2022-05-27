package handler

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/service"
	"github.com/Astronaut-X-X/TaskList/back_end/util"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user) // TODO deal err
	token, err, ok := service.LoginService(user)
	if ok {
		util.Response(c, http.StatusOK, 200, gin.H{"token": token})
	} else {
		util.Response(c, http.StatusBadRequest, 4000, err.Error())
	}
}

func RegisterHandler(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user) // TODO deal err
	res, ok := user.Create()
	if !ok {
		util.Response(c, http.StatusBadRequest, 4000, res.Error.Error())
	} else {
		util.Response(c, http.StatusOK, 200, util.ResMsg[200])
	}
}
