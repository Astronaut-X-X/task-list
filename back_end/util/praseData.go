package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Parse the JSON data in the request body.
func PraseData(c *gin.Context, object interface{}) {
	data, err := c.GetRawData()
	if err != nil {
		Response(c, http.StatusBadRequest, 4001, ResMsg[4001])
		c.Abort()
	}
	json.Unmarshal(data, &object)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
}
