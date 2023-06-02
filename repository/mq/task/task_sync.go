package task

import (
	"context"
	"encoding/json"

	"github.com/CocaineCong/micro-todoList/app/task/service"
	"github.com/CocaineCong/micro-todoList/consts"
	"github.com/CocaineCong/micro-todoList/idl"
	log "github.com/CocaineCong/micro-todoList/pkg/logger"
	"github.com/CocaineCong/micro-todoList/repository/mq"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskCreate(ctx context.Context) error {
	rabbitMqQueue := consts.RabbitMqTaskQueue
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return err
	}
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.LogrusObj.Infof("Received run Task: %s", d.Body)

			// 落库
			reqRabbitMQ := new(idl.TaskRequest)
			err = json.Unmarshal(d.Body, reqRabbitMQ)
			if err != nil {
				log.LogrusObj.Infof("Received run Task: %s", err)
			}

			err = service.TaskMQ2MySQL(ctx, reqRabbitMQ)
			if err != nil {
				log.LogrusObj.Infof("Received run Task: %s", err)
			}

		}
	}()

	log.LogrusObj.Infoln(err)
	<-forever

	return nil
}
