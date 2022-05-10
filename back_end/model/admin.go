// model 包定义了系统的模型
package model

import (
	"fmt"
	"github.com/Astronaut-X-X/GORM/config"
	"github.com/Astronaut-X-X/GORM/global"
	"github.com/Astronaut-X-X/GORM/utils"
)

//
//	InitAdminRole 初始化管理员用户 (初始化加载完毕后,以后都不会再执行)
//
func InitAdminRole(){
	result := global.GlobalDB.Where("Username = ?",config.ADMIN_NAME).Find(Admin{})
	utils.PrintResult(result)
	if result.RowsAffected == 0 {
		result = global.GlobalDB.Create(Admin{
			Username: config.ADMIN_NAME,
			Password: config.ADMIN_PASSWORD,
		})
		utils.PrintResult(result)
		if result.RowsAffected == 0{
			fmt.Println("Initializing admin role failed.")
		}else{
			fmt.Println("Initializing admin role succeed.")
		}
	}else{
		fmt.Println("Admin role already initialized")
	}
}

// Admin 定义了管理员用户 包括其中的基本信息
type Admin struct {
	 Username string `gorm:"not null;unique"` // 登录用户名
	 Password string `gorm:"not null"` // 登录密码
}

//
//	SelectAdminByUsername 根据用户名找到用户信息（用户名是唯一的）
//	参数：
//		username 用户名
//	返回值：
//		Admin 管理员用户对象
//
func SelectAdminByUsername(username string) Admin{
	var admin Admin
	result := global.GlobalDB.Where("Username = ?",username).Find(&admin)
	if result.Error != nil{
		// TODO 错误处理
	}
	return admin
}
