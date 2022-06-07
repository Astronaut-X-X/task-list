package validator

import (
	"regexp"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
)

func VerifyUserNamePasd(user model.User) (int, bool) {
	ok, err := regexp.MatchString(`[0-9a-zA-Z_\-]{6,16}`, user.Username)
	if !ok {
		return 4100, false
	}
	if err != nil {
		// TODO logging err
		return 5001, false
	}
	ok, err = regexp.MatchString(`[[0-9a-zA-Z@#$%^&*]{6,16}`, user.Password)
	if !ok {
		return 4101, false
	}
	if err != nil {
		// TODO logging err
		return 5001, false
	}
	return 3000, true
}

func VerifyUserEmail(user model.User) (int, bool) {
	ok, err := regexp.MatchString(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`, user.Email)
	if !ok {
		return 4102, false
	}
	if err != nil {
		// TODO logging err
		return 5001, false
	}
	return 3000, true
}
