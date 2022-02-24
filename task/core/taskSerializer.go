package core

import (
	"go_code/go-micro-todo/task/model"
	service "go_code/go-micro-todo/task/services"
)

func BuildTask(v model.Task) *service.TaskModel {
	return &service.TaskModel{
		Id: uint64(v.ID),
		Uid: uint64(v.Uid),
		Title: v.Title,
		Content: v.Content,
		Status: int64(v.Status),
		StartTime: v.StartTime,
		EndTime: v.EndTime,
	}
}
