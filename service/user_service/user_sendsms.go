package user_service

import (
	"context"
	"errors"
	"math/rand"
	"miaosha-jjl/common/proto/user"
	"miaosha-jjl/common/utils"
	"strconv"
)

func (s *ServerUser) SendSms(ctx context.Context, in *user.SendSmsRequest) (*user.SendSmsResponse, error) {
	code := rand.Intn(900) + 1000
	err := utils.SendSms(in.Mobile, strconv.Itoa(code))
	if err != nil {
		return nil, errors.New("验证码生成失败")
	}
	return &user.SendSmsResponse{Success: true}, nil
}
