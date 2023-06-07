package rpc

import (
	"context"

	"github.com/CocaineCong/micro-todoList/idl/pb"
	"github.com/CocaineCong/micro-todoList/pkg/e"
)

func TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	r, err := TaskService.CreateTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}

func TaskUpdate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	r, err := TaskService.UpdateTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}

func TaskDelete(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	r, err := TaskService.DeleteTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}

func TaskList(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskListResponse, err error) {
	r, err := TaskService.GetTasksList(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}
