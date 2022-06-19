package middleware

import (
	"fmt"
	"net/http"

	"github.com/Astronaut-X-X/TaskList/back_end/config"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			util.Response(c, http.StatusUnauthorized, 401, "Unauthorized")
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_MY_SECRET_KEY), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("%t", claims["id"])
			c.Set("id", claims["id"])
			c.Next()
		} else {
			util.Response(c, http.StatusBadRequest, 400, err.Error())
			c.Abort()
		}
	}
}
