package user_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
)

func (s *ServerUser) UserImproveInformation(ctx context.Context, in *user.UserImproveInformationRequest) (*user.UserImproveInformationResponse, error) {
	users := model.User{
		Uid:      uint32(in.Uid),
		Birthday: int32(in.Birthday),
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Addres:   in.Address,
	}
	err := users.UserImproveInformation(in.Uid)
	if err != nil {
		return nil, errors.New("用户完善信息失败")
	}
	return &user.UserImproveInformationResponse{Success: true}, nil
}
