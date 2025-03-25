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
func Update(ctx context.Context, i *user.UpdateRequest) (*user.UpdateResponse, error) {
	client, err := UserClient(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		login, err := client.Update(ctx, i)
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
