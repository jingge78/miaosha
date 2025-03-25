package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/product"
)

type ProductHandler func(ctx context.Context, client product.ProductClient) (interface{}, error)

func ProductClient(ctx context.Context, handler ProductHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := product.NewProductClient(dial)
	return handler(ctx, client)
}
func ProductDetail(ctx context.Context, in *product.ProductDetailRequest) (*product.ProductDetailResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.ProductDetail(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.ProductDetailResponse), nil
}
func EsAddProduct(ctx context.Context, in *product.EsAddProductRequest) (*product.EsAddProductResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.EsAddProduct(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.EsAddProductResponse), nil
}
func GetAllProduct(ctx context.Context, in *product.GetAllProductRequest) (*product.GetAllProductResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.GetAllProduct(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.GetAllProductResponse), nil
}
func EsSearchByKeyWord(ctx context.Context, in *product.EsSearchByKeyWordRequest) (*product.EsSearchByKeyWordResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.EsSearchByKeyWord(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.EsSearchByKeyWordResponse), nil
}
