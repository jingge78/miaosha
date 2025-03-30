package script

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"miaosha-jjl/common/global"
	"miaosha-jjl/common/model"
)

/*
   MysqlScript
   用来迁移数据库,只执行一次
*/

func MysqlScript() {
	var err error
	dsn := fmt.Sprintf("root:mysql_nNP5wS@tcp(117.27.231.169:3306)/seckill_project_db?charset=utf8mb4&parseTime=True&loc=Local")
	global.DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	global.DB.AutoMigrate(&model.UserMakeupCard{})

}
