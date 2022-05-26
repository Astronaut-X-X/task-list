package handler

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/util"

	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, nil)
}

func HomeHandler(c *gin.Context) {
	util.Response(c, http.StatusOK, 200, "OK")
}
