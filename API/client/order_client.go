package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/order"
)

type OrderHandler func(ctx context.Context, client order.OrderClient) (interface{}, error)

func OrderClient(ctx context.Context, handler OrderHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := order.NewOrderClient(dial)
	return handler(ctx, client)
}
func CreateOrder(ctx context.Context, i *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	client, err := OrderClient(ctx, func(ctx context.Context, client order.OrderClient) (interface{}, error) {
		createOrder, err := client.CreateOrder(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*order.CreateOrderResp), err
}
