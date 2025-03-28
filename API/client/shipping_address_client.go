package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/shipping_address"
)

type ShippingAddressHandler func(ctx context.Context, client shipping_address.ShippingAddressClient) (interface{}, error)

func ShippingAddressClient(ctx context.Context, handler ShippingAddressHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := shipping_address.NewShippingAddressClient(dial)
	return handler(ctx, client)
}
func AddShippingAddress(ctx context.Context, in *shipping_address.AddShippingAddressRequest) (*shipping_address.AddShippingAddressResponse, error) {
	client, err := ShippingAddressClient(ctx, func(ctx context.Context, client shipping_address.ShippingAddressClient) (interface{}, error) {
		createOrder, err := client.AddShippingAddress(ctx, in)
		if err != nil {
			return nil, err
		}
		return createOrder, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*shipping_address.AddShippingAddressResponse), err
}
