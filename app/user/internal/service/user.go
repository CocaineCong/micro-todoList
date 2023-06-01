package service

import (
	"context"
	"errors"
	"sync"

	"github.com/jinzhu/gorm"

	"github.com/CocaineCong/micro-todoList/app/user/internal/repository/db/dao"
	"github.com/CocaineCong/micro-todoList/app/user/internal/repository/db/model"
	"github.com/CocaineCong/micro-todoList/idl"
	"github.com/CocaineCong/micro-todoList/pkg/e"
)

func BuildUser(item *model.User) *idl.UserModel {
	userModel := idl.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *idl.UserRequest, resp *idl.UserDetailResponse) (err error) {
	resp = new(idl.UserDetailResponse)
	resp.Code = e.SUCCESS

	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		resp.Code = e.ERROR
		return
	}

	if !user.CheckPassword(req.Password) {
		resp.Code = e.InvalidParams
		return nil
	}

	resp.UserDetail = BuildUser(user)
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *idl.UserRequest, resp *idl.UserDetailResponse) (err error) {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	resp = new(idl.UserDetailResponse)
	resp.Code = e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	_, err = userDao.FindUserByUserName(req.UserName)
	if err != nil {
		if err == gorm.ErrRecordNotFound { // 如果不存在就继续下去
			// ...continue
		} else {
			resp.Code = e.ERROR
			return
		}
	}
	user := &model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = e.ERROR
		return err
	}
	if err = userDao.CreateUser(user); err != nil {
		resp.Code = e.ERROR
		return err
	}

	resp.UserDetail = BuildUser(user)
	return
}
