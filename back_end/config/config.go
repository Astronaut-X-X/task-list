package config

import (
	"fmt"
	"time"
)

// Database config
const DB_USERNAME string = "root"
const DB_PASSWORD string = "123456"
const DB_HOST string = "localhost"
const DB_PORT string = "3306"
const DB_DATABASE_NAME = "tl"
const DB_CHARSET = "utf8mb4"
const DB_PARSETIME = "true"  // Parse tiem
const DB_LOC = "Local"       // Local time
const DB_TablePrefix = "tl_" // Table Prefix

// Web config
const WEB_HOST string = "0.0.0.0"
const WEB_PORT string = "8081"
const WEB_RELATIVE_PATH = "/static"
const WEB_STATIC_ROOT = "./static"
const WEB_RUNMODE_DEBUG = true
const WEB_READTIMEOUT = 3 * 10 * time.Second
const WEB_WRITETIMEOUT = 3 * 10 * time.Second
const WEB_MAXHEADERBYTES = 1 << 20

// JWT
const JWT_MY_SECRET_KEY string = "307680186@qq.com"

func init() {
	fmt.Println("loading config ...")
}
