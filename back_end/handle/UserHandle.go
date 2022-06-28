package handle

import (
	"fmt"
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/service"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	v "github.com/Astronaut-X-X/TaskList/back_end/validator"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func LoginHandle(c *gin.Context) {
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

func RegisterHandle(c *gin.Context) {
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

func GetUserHandle(c *gin.Context) {
	id := c.GetFloat64("id")
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

func UploadUserImageHandle(c *gin.Context) {
	file, _ := c.FormFile("file")
	path := "./static"
	tmpFilename := uuid.New()
	dst := fmt.Sprintf("%s/%s.jpg", path, tmpFilename)
	filename := fmt.Sprintf("/%s.jpg", tmpFilename)
	user, ok := model.SelectUserById(int(c.GetFloat64("id")))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, gin.H{"msg": util.ResMsg[5002]})
	}
	user.Image = filename
	ok = user.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5003, gin.H{"msg": util.ResMsg[5003]})
	}
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 4000, gin.H{"msg": err.Error()})
	} else {
		util.Response(c, http.StatusOK, 200, gin.H{"filename": filename})
	}
}

func UpdateUserHandle(c *gin.Context) {
	var resUser model.User
	err := c.ShouldBindJSON(&resUser) // TODO deal err
	if err != nil {
		// TODO log
	}
	user, ok := model.SelectUserById(int(c.GetFloat64("id")))
	if !ok {
		util.Response(c, http.StatusBadRequest, 5002, gin.H{"msg": util.ResMsg[5002]})
	}
	user.Username = resUser.Username
	user.Email = resUser.Email
	ok = user.Update()
	if !ok {
		util.Response(c, http.StatusBadRequest, 5003, gin.H{"msg": util.ResMsg[5003]})
	}
	util.Response(c, http.StatusOK, 200, gin.H{"user": user})
}
