package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/micro-todoList/mq-server/model"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	return &TaskDao{NewDBClient(ctx)}
}

// ListTaskByUserId 获取user id
func (dao *TaskDao) ListTaskByUserId(userId uint64, start, limit int) (r []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).
		Limit(limit).Where("uid = ?", userId).
		Find(&r).Error

	err = dao.Model(&model.Task{}).Where("uid = ?", userId).
		Count(&count).Error

	return
}

// GetTaskByTaskIdAndUserId
func (dao *TaskDao) GetTaskByTaskIdAndUserId(taskId, userId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).
		Where("id = ? AND uid = ?", taskId, userId).
		First(&r).Error

	return
}
