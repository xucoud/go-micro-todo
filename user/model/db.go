package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_code/go-micro-todo/user/config"
)

var (
	Db *gorm.DB
	err error
)

func InitDB() {
	conStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Ms.DbUser,
		config.Ms.DbPassword,
		config.Ms.DbHost,
		config.Ms.DbPort,
		config.Ms.DbName)
	Db, err = gorm.Open(config.Ms.Db, conStr)
	if err != nil {
		fmt.Println("connect Mysql fail! err =", err)
		return
	}
	Db.AutoMigrate(&User{})
}
