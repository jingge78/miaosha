package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/initialition"
)

type IntegralogHandler func(ctx context.Context, client initialition.IntegrateClient) (interface{}, error)

func IntegralogClient(ctx context.Context, handler IntegralogHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := initialition.NewIntegrateClient(dial)
	return handler(ctx, client)
}
func Integration(ctx context.Context, in *initialition.IntegrationRequest) (*initialition.IntegrationResponse, error) {
	client, err := IntegralogClient(ctx, func(ctx context.Context, client initialition.IntegrateClient) (interface{}, error) {
		createOrder, err := client.Integration(ctx, in)
		if err != nil {
			return nil, err
		}
		return createOrder, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*initialition.IntegrationResponse), err
}
