package rpc

import (
	"github.com/micro/go-micro/v2"

	"github.com/CocaineCong/micro-todoList/app/gateway/wrappers"
	"github.com/CocaineCong/micro-todoList/idl"
)

var (
	UserService idl.UserService
	TaskService idl.TaskService
)

func InitRPC() {
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

	UserService = userService
	TaskService = taskService
}
