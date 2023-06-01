package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	"github.com/CocaineCong/micro-todoList/app/gateway/http"
	"github.com/CocaineCong/micro-todoList/app/gateway/wrappers"
	"github.com/CocaineCong/micro-todoList/idl"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	// 用户服务调用实例
	userService := idl.NewUserService("rpcUserService", userMicroService.Client())

	// task
	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	taskService := idl.NewTaskService("rpcTaskService", taskMicroService.Client())

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		// 将服务调用实例使用gin处理
		web.Handler(http.NewRouter(userService, taskService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
