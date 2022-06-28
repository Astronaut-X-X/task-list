package handle

import (
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/util"

	"github.com/gin-gonic/gin"
)

func DefaultHandle(c *gin.Context) {
	util.Response(c, http.StatusOK, 201, util.ResMsg[201])
}

func HomeHandle(c *gin.Context) {
	util.Response(c, http.StatusOK, 200, "OK")
}
