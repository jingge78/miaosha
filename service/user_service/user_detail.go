package user_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
)

func (s *ServerUser) UserDetail(ctx context.Context, in *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	var users model.User
	err := users.GetUserUid(in.Uid)
	if err != nil {
		return nil, err
	}
	if users.Uid == 0 {
		return nil, errors.New("用户不存在，请前往注册")
	}
	return &user.UserDetailResponse{
		Account:  users.Account,
		Birthday: uint64(users.Birthday),
		Nickname: users.Nickname,
		Avatar:   users.Avatar,
		Phone:    users.Phone,
		Address:  users.Addres,
	}, nil
}
