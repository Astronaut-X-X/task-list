package service

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/Astronaut-X-X/TaskList/back_end/config"
	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/Astronaut-X-X/TaskList/back_end/util"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	id uint
	jwt.Claims
}

var hmacSampleSecret = []byte(config.JWT_MY_SECRET_KEY)

func LoginService(user model.User) (string, error, bool) {
	user.Password = util.MD5Encipher(user.Password)
	tmpUser, ok := model.SelectUserByUsername(user.Username)
	if !ok {
		fmt.Println("User not")
		return "", errors.New("User is not exist"), false
	}
	if !reflect.DeepEqual(tmpUser.Password, user.Password) {
		fmt.Println("Pass")
		return "", errors.New("Password is not right"), false
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(360 * time.Hour).Unix(),
		"nbf": time.Now().Unix(),
		"id":  tmpUser.Model.ID,
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err, false
	}
	return tokenString, nil, true
}
