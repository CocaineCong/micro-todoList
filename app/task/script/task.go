package script

import (
	"context"

	log "github.com/CocaineCong/micro-todoList/pkg/logger"
	"github.com/CocaineCong/micro-todoList/repository/mq/task"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskCreate(ctx)
	if err != nil {
		log.LogrusObj.Infof("RunTaskCreate:%s", err)
	}
}
