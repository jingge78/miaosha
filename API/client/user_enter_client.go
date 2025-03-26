package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/user_enter"
)

type UserEnterHandler func(ctx context.Context, client user_enter.UserEnterClient) (interface{}, error)

func UserEnterClient(ctx context.Context, handler UserEnterHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := user_enter.NewUserEnterClient(dial)
	return handler(ctx, client)
}

func AddUserEnter(ctx context.Context, in *user_enter.AddUserEnterRequest) (*user_enter.AddUserEnterResponse, error) {
	client, err := UserEnterClient(ctx, func(ctx context.Context, client user_enter.UserEnterClient) (interface{}, error) {
		login, err := client.AddUserEnter(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user_enter.AddUserEnterResponse), nil
}
