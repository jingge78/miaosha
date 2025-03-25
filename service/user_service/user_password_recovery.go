package user_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
)

func (s *ServerUser) PasswordRecovery(ctx context.Context, in *user.PasswordRecoveryRequest) (*user.PasswordRecoveryResponse, error) {
	var users model.User
	err := users.LoginUser(in.Account)
	if err != nil {
		return nil, err
	}
	if users.Uid == 0 {
		return nil, errors.New("用户不存在，请前往注册")
	}

	return &user.PasswordRecoveryResponse{Password: users.Pwd}, err
}
