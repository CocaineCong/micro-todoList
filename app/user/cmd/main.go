package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/CocaineCong/micro-todoList/app/user/repository/db/dao"
	"github.com/CocaineCong/micro-todoList/app/user/service"
	"github.com/CocaineCong/micro-todoList/config"
	"github.com/CocaineCong/micro-todoList/idl/pb"
)

func main() {
	config.Init()
	dao.InitDB()
	// etcd注册件
	etcdReg := registry.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	// 启动微服务
	_ = microService.Run()
}
