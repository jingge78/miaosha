package user_service

import (
	"context"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user/user"
)

type ServerUser struct {
	user.UnimplementedUserServer
}

func (s ServerUser) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	var users model.User
	err := users.LoginUser(in.Account)
	if err != nil {
		return &user.LoginResponse{Result: "查询失败"}, err
	}
	if users.Uid == 0 {
		return &user.LoginResponse{Result: "用户不存在"}, err
	}
	if users.Pwd != in.Password {
		return &user.LoginResponse{Result: "密码错误"}, err
	}
	return &user.LoginResponse{Result: "登录成功"}, err
}
