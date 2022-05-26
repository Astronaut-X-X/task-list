package util

import "github.com/gin-gonic/gin"

var resMsg = map[int]string{
	200: "OK",
	201: "Empty message",
	404: "404",
}

func Response(c *gin.Context, httpCode int, code int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"data": data,
	})
}
