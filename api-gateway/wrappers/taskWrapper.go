package wrappers

import (
	"api-gateway/services"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)

func NewTask(id uint64, name string) *services.TaskModel {
	return &services.TaskModel{
		Id:         id,
		Title:      name,
		Content:    "响应超时",
		StartTime:  1000,
		EndTime:    1000,
		Status:     0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

// 降级函数
func DefaultTasks(resp interface{}) {
	models := make([]*services.TaskModel, 0)
	var i uint64
	for i := 0; i < 10; i++ {
		models = append(models, NewTask(i, "降级备忘录"+strconv.Itoa(20+int(i))))
	}
	result := resp.(*services.TaskListResponse)
	result.TaskList = models
}

type TaskWrapper struct {
	client.Client
}

func (wrapper *TaskWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                3000,
		RequestVolumeThreshold: 20,   //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		//降级处理
		//DefaultProducts(rsp)
		//DefaultData(rsp)
		return err
	})
}


//NewProductWrapper 初始化Wrapper
func NewTaskWrapper(c client.Client) client.Client {
	return &TaskWrapper{c}
}

