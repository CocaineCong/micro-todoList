package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/micro-todoList/app/gateway/rpc"
	"github.com/CocaineCong/micro-todoList/idl"
	"github.com/CocaineCong/micro-todoList/pkg/ctl"
	"github.com/CocaineCong/micro-todoList/pkg/utils"
	"github.com/CocaineCong/micro-todoList/types"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var req idl.UserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "UserRegister Bind 绑定参数失败"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserRegister RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var req idl.UserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "UserLogin Bind 绑定参数失败"))
		return
	}
	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserLogin RPC 调用失败"))
		return
	}
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "GenerateToken 失败"))
		return
	}
	res := &types.TokenData{
		User:  userResp,
		Token: token,
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
}
