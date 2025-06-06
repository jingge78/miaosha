package user_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
	"miaosha-jjl/common/utils"
)

// 修改密码//jj

func (s *ServerUser) Update(ctx context.Context, in *user.UpdateRequest) (*user.UpdateResponse, error) {
	var users model.User
	err := users.LoginUser(in.Account)
	if err != nil {
		return nil, err
	}
	if users.Uid == 0 {
		return nil, errors.New("账号不存在")
	}
	err = users.Update(in.Account, utils.Encryption(in.Password))
	if err != nil {
		return nil, err
	}
	return &user.UpdateResponse{Success: true}, nil
}
