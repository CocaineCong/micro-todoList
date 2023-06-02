package rpc

import (
	"context"

	"github.com/CocaineCong/micro-todoList/idl"
)

func TaskCreate(ctx context.Context, req *idl.TaskRequest) (resp *idl.TaskDetailResponse, err error) {
	r, err := TaskService.CreateTask(ctx, req)
	if err != nil {
		return
	}

	return r, nil
}

func TaskUpdate(ctx context.Context, req *idl.TaskRequest) (resp *idl.TaskDetailResponse, err error) {
	r, err := TaskService.UpdateTask(ctx, req)
	if err != nil {
		return
	}

	return r, nil
}

func TaskDelete(ctx context.Context, req *idl.TaskRequest) (resp *idl.TaskDetailResponse, err error) {
	r, err := TaskService.DeleteTask(ctx, req)
	if err != nil {
		return
	}

	return r, nil
}

func TaskList(ctx context.Context, req *idl.TaskRequest) (resp *idl.TaskListResponse, err error) {
	r, err := TaskService.GetTasksList(ctx, req)
	if err != nil {
		return
	}

	return r, nil
}
