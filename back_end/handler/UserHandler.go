package handler

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/service"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	v "github.com/Astronaut-X-X/TaskList/back_end/validator"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		// TODO log
		return
	}
	code, ok := v.VerifyUserNamePasd(user)
	if !ok {
		util.Response(c, http.StatusBadRequest, code, util.ResMsg[code])
		return
	}
	res, ok := service.LoginService(user)
	if ok {
		util.Response(c, http.StatusOK, 200, gin.H{"token": res})
	} else {
		util.Response(c, http.StatusBadRequest, 4000, res)
	}
}

func RegisterHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user) // TODO deal err
	if err != nil {
		// TODO log
	}
	code, ok := v.VerifyUserNamePasd(user)
	if !ok {
		util.Response(c, http.StatusBadRequest, code, util.ResMsg[code])
		return
	}
	code, ok = v.VerifyUserEmail(user)
	if !ok {
		util.Response(c, http.StatusBadRequest, code, util.ResMsg[code])
		return
	}
	res, ok := service.ReigsterUser(user)
	if ok {
		util.Response(c, http.StatusOK, 200, res)
	} else {
		util.Response(c, http.StatusBadRequest, 4000, res)
	}
}
