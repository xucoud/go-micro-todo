package main

import (
	"go_code/go-micro-todo/mq-server/conf"
	"go_code/go-micro-todo/mq-server/model"
	"go_code/go-micro-todo/mq-server/service"
)

func main() {
	conf.InitConfig()
	model.InitDB()
	model.InitMq()

	forever := make(chan bool)
	service.CreateTask()
	<- forever
}
