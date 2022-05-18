package service

import (
	"errors"
	"reflect"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	id uint
	jwt.Claims
}

var hmacSampleSecret = []byte{3, 0, 7, 6, 8, 0, 1, 8, 6}

func LoginService(username, password string) (token string, err error, ok bool) {

	if user, ok := model.SelectUserByUsername(username); ok {
		if reflect.DeepEqual(user.Password, password) {
			token := jwt.NewWithClaims(jwt.SigningMethodES256, &UserClaims{
				id: user.Model.ID,
			})
			tokenString, err := token.SignedString(hmacSampleSecret)
			if err != nil {
				return "", err, false
			}
			return tokenString, nil, true
		} else {
			return "", errors.New("Password is not right"), false
		}
	} else {
		return "", errors.New(("User is not exist")), false
	}

}
