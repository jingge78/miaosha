package user_enter_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user_enter"
)

func (e *ServerUserEnter) AddUserEnter(ctx context.Context, in *user_enter.AddUserEnterReq) (*user_enter.AddUserEnterResp, error) {
	userEnter := model.UserEnter{
		Uid:          uint32(in.Uid),
		Province:     in.Province,
		City:         in.City,
		District:     in.District,
		Address:      in.Address,
		MerchantName: in.Charter,
		LinkTel:      in.LinkTel,
		Charter:      in.Charter,
	}
	err := userEnter.AddUserEnter()
	if err != nil {
		return nil, err
	}
	if userEnter.Id == 0 {
		return nil, errors.New("未查询到商品信息")
	}
	return &user_enter.AddUserEnterResp{UserEnterId: uint64(userEnter.Id)}, nil
}
