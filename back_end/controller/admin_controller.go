package controller

import (
	"github.com/Astronaut-X-X/GORM/model"
	"regexp"
)

//
// Login 管理员用户 Admin 的登录业务处理函数
// 参数：
// 		username 用户名		（空判断，长度判断，合法字符）
// 		password 用户密码 	（空判断，长度判断，合法字符）
// 返回值：
// 		uint8 返回一个 uint8 值
//		- 登录成功		返回 1
//		- 用户不存在 	返回 2
//		- 用户密码错误 	返回 3
//		- 非法用户		返回 4
//		- 非法密码		返回 5
//
func Login(username,password  string) uint8 {
	if !CheckUsername(username) {
		return 4
	}
	if !CheckPassword(password){
		return 5
	}
	if admin := model.SelectAdminByUsername(username);admin.Username == ""{
		return 2
	}else if password != admin.Password{
		return 3
	}else{
		return 1
	}
}

//
// CheckUsername 用户名校验
// 参数：
// 		username 用户名（空判断，长度判断，合法字符）
// 返回值：
// 		bool 返回 true（符合） 和 false（不符合）
//
func CheckUsername(username string) bool{
	if len(username) == 0 || len(username) > 20 {
		return false
	}
	pattern := regexp.MustCompile("[0-9a-zA-Z_]{1:20}")
	return pattern.MatchString(username)
}

//
// CheckPassword 用户名校验
// 参数：
// 		password 用户名（空判断，长度判断，合法字符）
// 返回值：
// 		bool 返回 true（符合） 和 false（不符合）
//
func CheckPassword(password string) bool{
	if len(password) == 0 || len(password) > 20 {
		return false
	}
	pattern := regexp.MustCompile("[0-9a-zA-Z@#$%^&*]{1:20}")
	return pattern.MatchString(password)
}
