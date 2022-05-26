package main

import (
	"github.com/Astronaut-X-X/TaskList/back_end/model"
)

func main() {

	user := new(model.User)
	user.Model.ID = 1
	user.Username = "CCSSD"
	user.Password = "123456"
	user.Update()

}
