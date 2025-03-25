package user_service

<<<<<<< HEAD
import (
	"context"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user/user"
)
=======
import "miaosha-jjl/common/proto/user"
>>>>>>> f7cd67524df94e75b5a629c18176469208cedc32

type ServerUser struct {
	user.UnimplementedUserServer
}
<<<<<<< HEAD

func (s ServerUser) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	var users model.User
	err := users.LoginUser(in.Account)
	if err != nil {
		return &user.LoginResponse{Result: "查询失败"}, err
	}
	if users.Uid == 0 {
		return &user.LoginResponse{Result: "用户不存在"}, err
	}
	if users.Pwd != in.Password {
		return &user.LoginResponse{Result: "密码错误"}, err
	}
	return &user.LoginResponse{Result: "登录成功"}, err
}
=======
>>>>>>> f7cd67524df94e75b5a629c18176469208cedc32
