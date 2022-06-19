package service

import (
	"reflect"
	"time"

	"github.com/Astronaut-X-X/TaskList/back_end/config"
	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = []byte(config.JWT_MY_SECRET_KEY)

func LoginService(user model.User) (string, bool) {
	user.Password = util.MD5Encipher(user.Password)
	tmpUser, ok := model.SelectUserByUsername(user.Username)
	if !ok {
		return "User is not exist", false
	}
	if !reflect.DeepEqual(tmpUser.Password, user.Password) {
		return "Password is not right", false
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"nbf": time.Now().Unix(),
		"id":  tmpUser.Model.ID,
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		// TODO log recode err
		return "Token generation failed", false
	}
	return tokenString, true
}

func ReigsterUser(user model.User) (string, bool) {
	_, ok := model.SelectUserByUsername(user.Username)
	if ok {
		return "User is exist", false
	}
	user.Password = util.MD5Encipher(user.Password)
	if user.Create() {
		return "User generation succeeded", true
	}
	return "Error db option", false
}

func FlashToken(tokenStr string) (string, bool) {
	if tokenStr == "" {
		return "Token is empty", false
	}
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_MY_SECRET_KEY), nil
	})
	if err != nil {
		return "Token prase error", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(float64)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"nbf": time.Now().Unix(),
			"id":  id,
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		if err != nil {
			return "Token generate error", false
		}
		return tokenString, true
	} else {
		return "Token lose efficacy", false
	}
}

func UpdateUser(){
	
}