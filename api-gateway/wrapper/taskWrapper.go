package wrapper

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	service "go_code/go-micro-todo/api-gateway/services"
	"strconv"
)

func NewTask(id uint64, name string) *service.TaskModel {
	return &service.TaskModel{
		Id: id,
		Title: name,
		Content: "响应超时",
		Status: 0,
		StartTime: 1000,
		EndTime: 1000,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

//降级函数
func DefaultTasks(rep interface{}) {
	models := make([]*service.TaskModel, 0)
	for i := 0; i < 10; i++ {
		models = append(models, NewTask(uint64(i), "降级备忘录"+strconv.Itoa(i)))
	}
	result := rep.(service.TaskListResponse)
	result.TaskList = models
}

type taskWrapper struct {
	client.Client
}

func (wrapper *taskWrapper) Call(ctx context.Context,req client.Request,resp interface{},opts ...client.CallOption)error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,   //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

func NewTaskWrapper(c client.Client) client.Client {
	return &taskWrapper{c}
}
