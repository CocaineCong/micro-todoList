package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/CocaineCong/micro-todoList/app/gateway/rpc"
	"github.com/CocaineCong/micro-todoList/idl/pb"
	"github.com/CocaineCong/micro-todoList/pkg/ctl"
)

func ListTaskHandler(ctx *gin.Context) {
	var taskReq pb.TaskRequest
	if err := ctx.Bind(&taskReq); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "获取用户信息错误"))
		return
	}
	taskReq.Uid = uint64(user.Id)
	// 调用服务端的函数
	taskResp, err := rpc.TaskList(ctx, &taskReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "taskResp RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskResp))
}

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "获取用户信息错误"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskCreate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskList RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func GetTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "获取用户信息错误"))
		return
	}
	req.Id = cast.ToUint64(ctx.Param("id"))
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskGet(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskList RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "获取用户信息错误"))
		return
	}
	req.Id = cast.ToUint64(ctx.Param("id"))
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskUpdate RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func DeleteTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "获取用户信息错误"))
		return
	}
	req.Id = cast.ToUint64(ctx.Param("id"))
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskDelete RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}
