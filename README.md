# micro-todoList
# Go-Micro V2 + RabbitMQ 构造简单备忘录

本项目改自于作者[Congz](https://github.com/congz666)的[后台管理系统](https://github.com/congz666/backstage-go)

将原项目的micro升到v2，服务发现也换成etcd。

保留原作者的熔断机制，token验证，网关和各模块之间的rpc通信等

在此也非常感谢作者开源！

# 项目的详细博客地址

**[用户模块](https://blog.csdn.net/weixin_45304503/article/details/122286980)**

**[备忘录模块](https://blog.csdn.net/weixin_45304503/article/details/122301707)**

# 项目的视频介绍地址

[Go-Micro+RabbitMQ 构建简单备忘录](https://www.bilibili.com/video/BV1h44y1L7LN)

# 项目的主要功能介绍

- 用户注册登录 ( jwt-go鉴权 )
- 新增/删除/修改/查询 备忘录

# 项目主要依赖：

**Golang V1.16**

- Gin
- Gorm
- mysql
- go-micro
- protobuf
- grpc
- amqp
- ini
- hystrix
- jwt-go
- crypto

# 项目结构

## 1. gateway 网关部分

```
gateway/
├── pkg
│  ├── e
│  ├── logging
│  ├── util
├── services
│  ├── proto
├── weblib
│  ├── handlers
│  ├── middleware
└── wrappers
```
- pkg/e : 封装错误码
- pkg/logging : 日志文件
- pkg/util : 工具函数
- service/proto : 放置proto文件以及生成的pb文件
- weblib/handlers : 各个服务的接口
- weblib/middleware : http服务器的中间件
- wrappers : 放置服务熔断的配置

## 2. mq-server RabbitMQ 消息队列

```
mq-server/
├── conf
├── model
└── service
```

- conf：配置信息
- model：数据库模型
- service：服务

## 3. task & user

```
task/ & user/
├── conf
├── core
├── model
└── service
```

- conf：配置信息
- core：业务逻辑
- model：数据库模型
- service：proto文件以及各服务

conf/config.ini 文件
```ini
[service]
AppMode = debug
HttpPort = :3000

[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = root
DbName = micro_todolist

[rabbitmq]
RabbitMQ = amqp
RabbitMQUser = guest
RabbitMQPassWord = guest
RabbitMQHost = localhost
RabbitMQPort = 5672
```

# 运行简要说明
1. 保证rabbitMQ开启状态
2. 保证etcd开启状态
3. 依次执行各模块下的main.go文件
4. 执行user,task,api-gateway的时候需要后面加上这个，注册到etcd并且注册地址是这个地址。
```go
go run main.go --registry=etcd --registry_address=127.0.0.1:2379
```

**如果出错一定要注意打开etcd的keeper查看服务是否注册到etcd中。**


# 最后

再次感谢原作者[Congz](https://github.com/congz666)的开源




