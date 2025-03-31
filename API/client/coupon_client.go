package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/coupon"
)

type CouponHandler func(ctx context.Context, client coupon.CouponClient) (interface{}, error)

func CouponClient(ctx context.Context, handler CouponHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := coupon.NewCouponClient(dial)
	return handler(ctx, client)
}

func AddCoupon(ctx context.Context, in *coupon.AddCouponRequest) (*coupon.AddCouponResponse, error) {
	client, err := CouponClient(ctx, func(ctx context.Context, client coupon.CouponClient) (interface{}, error) {
		createOrder, err := client.AddCoupon(ctx, in)
		if err != nil {
			return nil, err
		}
		return createOrder, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*coupon.AddCouponResponse), err
}
