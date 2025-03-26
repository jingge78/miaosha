package user_enter_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user_enter"
)

type ServerUserEnter struct {
	user_enter.UnimplementedUserEnterServer
}

func (e *ServerUserEnter) AddUserEnter(ctx context.Context, in *user_enter.AddUserEnterRequest) (*user_enter.AddUserEnterResponse, error) {
	// 判断用户是否已经注册商户
	userEnter := model.UserEnter{}
	err := userEnter.GetUserEnterId(in.Uid)
	if err != nil {
		return nil, errors.New("查询商户失败")
	}
	if userEnter.Id != 0 {
		return nil, errors.New("该用户已经注册商户")
	}
	// 注册商户
	userEnter = model.UserEnter{
		Uid:          uint32(in.Uid),
		Province:     in.Province,
		City:         in.City,
		District:     in.District,
		Address:      in.Address,
		MerchantName: in.MerchantName,
		LinkTel:      in.LinkTel,
		Charter:      in.Charter,
	}
	err = userEnter.AddUserEnter()
	if err != nil {
		return nil, err
	}
	if userEnter.Id == 0 {
		return nil, errors.New("商户入驻失败")
	}
	return &user_enter.AddUserEnterResponse{UserEnterId: uint64(userEnter.Id)}, nil
}
