package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go_code/go-micro-todo/task/conf"
	"go_code/go-micro-todo/task/core"
	"go_code/go-micro-todo/task/model"
	service "go_code/go-micro-todo/task/services"
)

func main() {
	conf.InitConfig()
	model.InitDB()
	model.InitMq()

	//注册etcd
	etcdResg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		)

	//注册微服务实例
	taskMicro := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdResg),
		)
	//初始话
	taskMicro.Init()
	//服务注册
	_ = service.RegisterTaskServiceHandler(taskMicro.Server(), new(core.TaskService))
	//启动服务
	_ = taskMicro.Run()
}
