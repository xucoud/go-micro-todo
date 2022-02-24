package service

import (
	"encoding/json"
	"go_code/go-micro-todo/mq-server/model"
	"log"
)

//从RabbitMQ中读取消息
func CreateTask() {
	ch, err := model.Mq.Channel()
	if err != nil {
		panic(err)
	}
	q, _ := ch.QueueDeclare("task_mq", true, false, false,false, nil)
	err = ch.Qos(1, 0, false)
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume(q.Name, "", false,false,false,false,nil)
	if err != nil {
		panic(err)
	}
	//处于监听状态，一直监听生产端的生产，阻塞主线程
	go func() {
		for d := range msgs {
			var t model.Task
			err = json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.Db.Create(&t)
			log.Println("create task done")
			_ = d.Ack(false)
		}
	}()
}
