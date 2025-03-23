package user_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
	"miaosha-jjl/common/utils"
)

func (s *ServerUser) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	var users model.User
	err := users.LoginUser(in.Account)
	if err != nil {
		return nil, err
	}
	if users.Uid != 0 {
		return nil, errors.New("用户已存在，请前往登录")
	}
	userAdd := model.User{
		Account: in.Account,
		Pwd:     utils.Md5(in.Password),
	}
	err = userAdd.Create()
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{Success: true}, nil
}
