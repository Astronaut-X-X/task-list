package validator

import (
	"regexp"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/gin-gonic/gin"
)

func VerifyUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		util.PraseData(c, user)
		if ok, _ := regexp.MatchString(``, user.Username); !ok {

		}
	}
}
