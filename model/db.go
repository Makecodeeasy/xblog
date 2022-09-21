package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"xblog/utils"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("数据库连接失败，请检查参数：%s", err))
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("查询数据库失败: %s", err))
	}
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	// 关闭
	// sqlDB.Close()
}
