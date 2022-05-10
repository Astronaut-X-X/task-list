package global

import (
	"database/sql"
	"gorm.io/gorm"
)

// 存放一个 Gorm 的全局 DB 为任何地方使用
var GlobalDB *gorm.DB
// 存放一个 sql 的全局 DB
var SqlDB *sql.DB

