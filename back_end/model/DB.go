package model

import (
	"fmt"

	"github.com/Astronaut-X-X/TaskList/back_end/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {

	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE_NAME,
		config.DB_CHARSET,
		config.DB_PARSETIME,
		config.DB_LOC,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.DB_TablePrefix,
		},
	})
	// Debug
	DB = DB.Debug()

	if err != nil {
		// TODO log
		panic("Something worry with open db.")
	}
	migrate()

}

func migrate() {
	DB.AutoMigrate(
		&User{},
		&Task{},
		&Goal{},
		&Todo{},
		&DailyPlan{},
		&DailyDetail{},
		&WeekPlan{},
		&Happy{},
	)
}
