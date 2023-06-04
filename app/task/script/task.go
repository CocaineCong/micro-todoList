package script

import (
	"context"

	"github.com/CocaineCong/micro-todoList/app/task/repository/mq/task"
	log "github.com/CocaineCong/micro-todoList/pkg/logger"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskCreate(ctx)
	if err != nil {
		log.LogrusObj.Infof("RunTaskCreate:%s", err)
	}
}
