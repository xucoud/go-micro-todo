package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"go_code/go-micro-todo/task/model"
	service "go_code/go-micro-todo/task/services"
)

func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, rep *service.TaskDetailResponse) error {
	ch, err := model.Mq.Channel()
	if err != nil {
		err = errors.New("rabbitMQ channel fail, err =" + err.Error())
	}
	q, _ := ch.QueueDeclare("task_mq", true, false, false,false, nil)
	body, _ := json.Marshal(req)
	err = ch.Publish("", q.Name, false,false,amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "application/json",
		Body: body,
	})
	if err != nil {
		err = errors.New("rabbitMQ Publish fail, err =" + err.Error())
	}
	return nil
}

func (*TaskService) GetTasksList(ctx context.Context, req *service.TaskRequest, rep *service.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []*model.Task
	var count uint32
	//查找备忘录
	err := model.Db.Where("uid = ?", req.Uid).Find(&taskData).Error
	if err != nil {
		return errors.New("mysql find by uid fail, err =" + err.Error())
	}
	//统计数量
	model.Db.Model(model.Task{}).Where("uid = ?", req.Uid).Count(&count)
	//将结果转化为proto里分数据类型
	var taskRes []*service.TaskModel
	for _, v := range taskData {
		taskRes = append(taskRes, BuildTask(*v))
	}
	rep.TaskList = taskRes
	rep.Count = count
	return nil
}

func (*TaskService) GetTask(ctx context.Context, req *service.TaskRequest, rep *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.Db.First(&taskData, req.Id)
	rep.TaskDetail = BuildTask(taskData)
	return nil
}

func (*TaskService) UpdateTask(ctx context.Context, req *service.TaskRequest, rep *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.Db.Model(model.Task{}).Where("id = ? & uid = ?", req.Id, req.Uid).First(&taskData)
	taskData.Status = int(req.Status)
	taskData.Title = req.Title
	taskData.Content = req.Content
	model.Db.Save(&taskData)
	rep.TaskDetail = BuildTask(taskData)
	return nil
}

func (*TaskService) DeleteTask(ctx context.Context, req *service.TaskRequest, rep *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.Db.Model(model.Task{}).Where("id = ? & uid = ?", req.Id, req.Uid).First(&taskData)
	model.Db.Delete(&taskData)
	rep.TaskDetail = BuildTask(taskData)
	return nil
}
