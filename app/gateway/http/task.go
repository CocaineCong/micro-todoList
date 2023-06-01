package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/micro-todoList/idl"
	"github.com/CocaineCong/micro-todoList/pkg/utils"
)

func GetTaskList(ctx *gin.Context) {
	var taskReq idl.TaskRequest
	if err := ctx.Bind(&taskReq); err != nil {
		ctx.JSON(http.StatusBadRequest, "")
	}
	// 调用服务端的函数
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"task":  taskResp.TaskList,
			"count": taskResp.Count,
		},
	})
}

func CreateTask(ctx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.keys取出服务实例
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskService := ctx.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{"data": taskRes.TaskDetail})
}

func GetTaskDetail(ctx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ctx.BindUri(&taskReq))
	// 从gin.keys取出服务实例
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ctx.Param("id")) // 获取task_id
	taskReq.Id = uint64(id)
	productService := ctx.Keys["taskService"].(services.TaskService)
	productRes, err := productService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{"data": productRes.TaskDetail})
}

func UpdateTask(ctx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.keys取出服务实例
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	id, _ := strconv.Atoi(ctx.Param("id"))
	taskReq.Id = uint64(id)
	taskReq.Uid = uint64(claim.Id)
	taskService := ctx.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{"data": taskRes.TaskDetail})
}

func DeleteTask(ctx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.keys取出服务实例
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ctx.Param("id"))
	taskReq.Id = uint64(id)
	taskService := ctx.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{"data": taskRes.TaskDetail})
}
