package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	ser "go_code/go-micro-todo/api-gateway/services"
	"go_code/go-micro-todo/api-gateway/weblib"
	"go_code/go-micro-todo/api-gateway/wrapper"
	"time"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		)

	service := micro.NewService(
		micro.Name("userService.Client"),
		micro.WrapClient(wrapper.NewUserWrapper),
		)

	task := micro.NewService(
		micro.Name("taskService.Client"),
		micro.WrapClient(wrapper.NewTaskWrapper),
		)

	userService := ser.NewUserService("rpcUserService", service.Client())
	taskService := ser.NewTaskService("rpcTaskService", task.Client())

	server := web.NewService(
		web.Name("httpService"),
		web.Address(":4000"),
		web.Handler(weblib.NewRouter(userService, taskService)),
		web.Registry(etcdReg),
		web.RegisterTTL(30*time.Second),
		web.RegisterInterval(15*time.Second),
		web.Metadata(map[string]string{"protocol": "http"}),
		)

	_ = server.Init()

	_ = server.Run()
}