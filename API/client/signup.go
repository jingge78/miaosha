package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/signup"
)

type SignupHandler func(ctx context.Context, client signup.SignupClient) (interface{}, error)

func SignupClient(ctx context.Context, handler SignupHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := signup.NewSignupClient(dial)
	return handler(ctx, client)
}
func Send(ctx context.Context, i *signup.SignupRequest) (*signup.SignupResponse, error) {
	client, err := SignupClient(ctx, func(ctx context.Context, client signup.SignupClient) (interface{}, error) {
		createOrder, err := client.Send(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*signup.SignupResponse), err
}
