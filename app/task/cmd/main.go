package main

import (
	"context"

	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/CocaineCong/micro-todoList/app/task/repository/db/dao"
	"github.com/CocaineCong/micro-todoList/app/task/repository/mq"
	"github.com/CocaineCong/micro-todoList/app/task/script"
	"github.com/CocaineCong/micro-todoList/app/task/service"
	"github.com/CocaineCong/micro-todoList/config"
	"github.com/CocaineCong/micro-todoList/idl/pb"
	log "github.com/CocaineCong/micro-todoList/pkg/logger"
)

func main() {
	config.Init()
	dao.InitDB()
	mq.InitRabbitMQ()
	log.InitLog()

	// 启动一些脚本
	loadingScript()

	// etcd注册件
	etcdReg := registry.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"), // 微服务名字
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())
	// 启动微服务
	_ = microService.Run()
}

func loadingScript() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
