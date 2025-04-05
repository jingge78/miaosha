package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/product_reply"
)

type ProductReplyHandler func(ctx context.Context, client product_reply.ProductReplyClient) (interface{}, error)

func ProductReplyClient(ctx context.Context, handler ProductReplyHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := product_reply.NewProductReplyClient(dial)
	return handler(ctx, client)
}
func ProductReplyList(ctx context.Context, in *product_reply.ProductReplyListRequest) (*product_reply.ProductReplyListResponse, error) {
	client, err := ProductReplyClient(ctx, func(ctx context.Context, client product_reply.ProductReplyClient) (interface{}, error) {
		list, err := client.ProductReplyList(ctx, in)
		if err != nil {
			return nil, err
		}
		return list, nil

	})

	if err != nil {
		return nil, err
	}
	return client.(*product_reply.ProductReplyListResponse), nil
}
