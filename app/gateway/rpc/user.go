package rpc

import (
	"context"
	"errors"

	"github.com/CocaineCong/micro-todoList/idl"
	"github.com/CocaineCong/micro-todoList/pkg/e"
)

// UserLogin 用户登陆
func UserLogin(ctx context.Context, req *idl.UserRequest) (resp *idl.UserDetailResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil {
		return
	}
	if resp.Code != e.SUCCESS {
		err = errors.New("UserLogin 出现错误")
		return
	}

	return
}

// UserRegister 用户注册
func UserRegister(ctx context.Context, req *idl.UserRequest) (resp *idl.UserDetailResponse, err error) {
	resp, err = UserService.UserRegister(ctx, req)
	if err != nil {
		return
	}
	if resp.Code != e.SUCCESS {
		err = errors.New("UserRegister 出现错误")
		return
	}

	return
}
