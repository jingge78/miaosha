package signup_service

import (
	"context"
	"fmt"
	"log"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/signup"
	"miaosha-jjl/common/utils"
)

type ServerSignups struct {
	signup.UnimplementedSignupServer
}

func (s *ServerSignups) Send(ctx context.Context, in *signup.SignupRequest) (*signup.SignupResponse, error) {
	fmt.Println("000000")
	signupes := model.Email{
		Emaile:   in.Email,
		Password: in.Password,
	}
	fmt.Println("1111111111111111111111111")
	err := signupes.RegisterUser()
	if err != nil {
		return nil, err
	}
	fmt.Println("222222222222222")
	if err != nil {
		return nil, err
	}
	fmt.Printf("4444444444444444444")
	if signupes.Uid == 0 {
		return nil, err
	}
	fmt.Println("3333333333333333333")
	// 发送验证邮件
	utils.SendEmail(in.Email)
	fmt.Println("66666666")

	if err != nil {
		log.Fatalf("验证邮件发送失败: %v", err)

	}
	fmt.Println("555555555555555555")

	fmt.Println("用户注册成功，验证邮件已发送！")
	return &signup.SignupResponse{Result: int64(signupes.Uid)}, nil
}
