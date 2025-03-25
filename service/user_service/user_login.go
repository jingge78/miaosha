package user_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
	"miaosha-jjl/common/utils"
)

type ServerUser struct {
	user.UnimplementedUserServer
}

// 账号登录//jj

func (s *ServerUser) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	var users model.User
	err := users.LoginUser(in.Account)
	if err != nil {
		return nil, err
	}
	if users.Uid == 0 {
		return nil, errors.New("用户不存在，请前往注册")
	}
	if users.Pwd != utils.Encryption(in.Password) {
		return nil, errors.New("密码错误")
	}
	return &user.LoginResponse{UserId: int64(users.Uid)}, err
}
