package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go_code/go-micro-todo/user/config"
	"go_code/go-micro-todo/user/core"
	"go_code/go-micro-todo/user/model"
	"go_code/go-micro-todo/user/services"
)

func main() {
	config.InitConfig()
	model.InitDB()
	defer model.Db.Close()

	//etcd注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		)
	//微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),
		)

	microService.Init()

	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))

	_ = microService.Run()
}