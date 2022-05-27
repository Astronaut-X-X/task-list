package util

import "github.com/gin-gonic/gin"

var ResMsg = map[int]string{
	200:  "OK",
	201:  "Empty message",
	404:  "404",
	4000: "",
	4001: "Parsing JSON failed",
	4100: "The user name format is incorrect",
	4101: "The user password format is incorrect",
	4102: "The user email format is incorrect",
}

func Response(c *gin.Context, httpCode int, code int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"data": data,
	})
}
