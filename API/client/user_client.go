package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/user"
)

type UserHandler func(ctx context.Context, client user.UserClient) (interface{}, error)

func UserClient(ctx context.Context, handler UserHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := user.NewUserClient(dial)
	return handler(ctx, client)
}
func UserLogin(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		login, err := client.Login(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.LoginResponse), nil
}
func UserRegister(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		login, err := client.Register(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.RegisterResponse), nil
}
func Update(ctx context.Context, in *user.UpdateRequest) (*user.UpdateResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		login, err := client.Update(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UpdateResponse), nil
}

func SendSms(ctx context.Context, in *user.SendSmsRequest) (*user.SendSmsResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		login, err := client.SendSms(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.SendSmsResponse), nil
}

func PasswordRecovery(ctx context.Context, in *user.PasswordRecoveryRequest) (*user.PasswordRecoveryResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		pass, err := client.PasswordRecovery(ctx, in)
		if err != nil {
			return nil, err
		}
		return pass, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.PasswordRecoveryResponse), nil
}

func UserSignIn(ctx context.Context, in *user.SignInRequest) (*user.SignInResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		pass, err := client.SignIn(ctx, in)
		if err != nil {
			return nil, err
		}
		return pass, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.SignInResponse), nil
}

func UserDetail(ctx context.Context, in *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		detail, err := client.UserDetail(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}

	return client.(*user.UserDetailResponse), nil
}
func MakeupSignIn(ctx context.Context, in *user.MakeupSignInRequest) (*user.MakeupSignInResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		pass, err := client.MakeupSignIn(ctx, in)
		if err != nil {
			return nil, err
		}
		return pass, nil
	})
	if err != nil {
		return nil, err
	}

	return client.(*user.MakeupSignInResponse), nil
}
func UserImproveInformation(ctx context.Context, in *user.UserImproveInformationRequest) (*user.UserImproveInformationResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		improve, err := client.UserImproveInformation(ctx, in)
		if err != nil {
			return nil, err
		}
		return improve, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UserImproveInformationResponse), nil
}
