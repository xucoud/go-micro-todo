package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
	"go_code/go-micro-todo/task/conf"
	"strings"
	"time"
)

var (
	Db *gorm.DB
	Mq *amqp.Connection
	err error
)

func InitDB() {
	conStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		conf.Ms.DbUser,
		conf.Ms.DbPassword,
		conf.Ms.DbHost,
		conf.Ms.DbPort,
		conf.Ms.DbName)
	Db, err = gorm.Open(conf.Ms.Db, conStr)
	if err != nil {
		fmt.Println("connect Mysql fail! err =", err)
		return
	}
	Db.LogMode(true)
	if gin.Mode() == "release" {
		Db.LogMode(false)
	}
	Db.DB().SetMaxIdleConns(20)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(30 * time.Second)

	Db.AutoMigrate(&Task{})
}

func InitMq() {

	//链接MQ
	MQPath := strings.Join([]string{conf.MQ.Mq, "://", conf.MQ.User, ":", conf.MQ.Password, "@", conf.MQ.Host, ":", conf.MQ.Port, "/"}, "")
	Mq, err = amqp.Dial(MQPath)
	if err != nil {
		panic(err.Error())
	}

}
