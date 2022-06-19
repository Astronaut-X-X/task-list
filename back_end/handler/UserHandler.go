package handler

import (
	"fmt"
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
	// code, ok = v.VerifyUserEmail(user)
	// if !ok {
	// 	util.Response(c, http.StatusBadRequest, code, util.ResMsg[code])
	// 	return
	// }
	res, ok := service.ReigsterUser(user)
	if ok {
		util.Response(c, http.StatusOK, 200, res)
	} else {
		util.Response(c, http.StatusBadRequest, 4000, res)
	}
}

func GetUserByIDInContext(c *gin.Context) {
	id := c.GetFloat64("id")
	fmt.Println(id)
	user, ok := model.SelectUserById(int(id))
	if ok {
		user.Password = ""
		util.Response(c, http.StatusOK, 200, user)
	} else {
		util.Response(c, http.StatusBadRequest, 5002, util.ResMsg[5002])
	}
}

func FlashToken(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	res, ok := service.FlashToken(tokenStr)
	if ok {
		util.Response(c, http.StatusOK, 200, gin.H{"token": res})
	} else {
		util.Response(c, http.StatusBadRequest, 4000, gin.H{"msg": res})
	}
}
