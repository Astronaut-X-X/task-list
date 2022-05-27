package handler

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/util"

	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	util.Response(c, http.StatusOK, 201, util.ResMsg[201])
}

func HomeHandler(c *gin.Context) {
	util.Response(c, http.StatusOK, 200, "OK")
}
