package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/micro-todoList/idl"
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

// GetTaskByTaskIdAndUserId 通过 task id 和 user id 获取task
func (dao *TaskDao) GetTaskByTaskIdAndUserId(taskId, userId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).
		Where("id = ? AND uid = ?", taskId, userId).
		First(&r).Error

	return
}

// UpdateTask 更新task
func (dao *TaskDao) UpdateTask(req *idl.TaskRequest) (r *model.Task, err error) {
	taskData := new(model.Task)
	err = dao.Model(&model.Task{}).
		Where("id= ? AND uid=?", req.Id, req.Uid).
		First(&taskData).Error
	if err != nil {
		return
	}
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content

	err = dao.Model(&model.Task{}).Save(&taskData).Error
	return taskData, err
}

func (dao *TaskDao) DeleteTaskByIdAndUserId(id, uId uint64) error {
	return dao.Model(&model.Task{}).
		Where("id =? AND uid=?", id, uId).
		Delete(&model.Task{}).Error

}
