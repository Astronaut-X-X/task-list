package validator

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/Astronaut-X-X/TaskList/back_end/config"
	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"
	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func VerifyUserNamePasd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		util.PraseData(c, &user)
		fmt.Println(user)
		if ok, err := regexp.MatchString(`[0-9a-zA-Z_\-]{6,16}`, user.Username); !ok || err != nil {
			// TODO logging err
			fmt.Println(err)
			util.Response(c, http.StatusBadRequest, 4100, util.ResMsg[4100])
			c.Abort()
		} else if ok, err := regexp.MatchString(`[[0-9a-zA-Z@#$%^&*]{6,16}`, user.Password); !ok || err != nil {
			// TODO logging err
			fmt.Println(err)
			util.Response(c, http.StatusBadRequest, 4101, util.ResMsg[4101])
			c.Abort()
		}
		c.Next()
	}
}

func VerifyUserEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		util.PraseData(c, &user)
		if ok, err := regexp.MatchString(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`, user.Email); !ok || err != nil {
			// TODO logging err
			util.Response(c, http.StatusBadRequest, 4102, util.ResMsg[4102])
			c.Abort()
		}
		c.Next()
	}
}

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
			c.Set("id", claims["id"])
			c.Next()
		} else {
			util.Response(c, http.StatusBadRequest, 400, err.Error())
			c.Abort()
		}
	}
}
