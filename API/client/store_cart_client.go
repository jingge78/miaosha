package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/store_cart"
)

type StoreCartHandler func(ctx context.Context, client store_cart.StoreCartClient) (interface{}, error)

func StoreCartClient(ctx context.Context, handler StoreCartHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := store_cart.NewStoreCartClient(dial)
	return handler(ctx, client)
}
func AddStoreCart(ctx context.Context, in *store_cart.AddStoreCartRequest) (*store_cart.AddStoreCartResponse, error) {
	client, err := StoreCartClient(ctx, func(ctx context.Context, client store_cart.StoreCartClient) (interface{}, error) {
		cart, err := client.AddStoreCart(ctx, in)
		if err != nil {
			return nil, err
		}
		return cart, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*store_cart.AddStoreCartResponse), nil
}

func DeleteCart(ctx context.Context, in *store_cart.DeleteCartRequest) (*store_cart.DeleteCartResponse, error) {
	client, err := StoreCartClient(ctx, func(ctx context.Context, client store_cart.StoreCartClient) (interface{}, error) {
		del, err := client.DeleteCart(ctx, in)
		if err != nil {
			return nil, err
		}
		return del, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*store_cart.DeleteCartResponse), nil
}

func ClearCart(ctx context.Context, in *store_cart.ClearCartRequest) (*store_cart.ClearCartResponse, error) {
	client, err := StoreCartClient(ctx, func(ctx context.Context, client store_cart.StoreCartClient) (interface{}, error) {
		clears, err := client.ClearCart(ctx, in)
		if err != nil {
			return nil, err
		}
		return clears, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*store_cart.ClearCartResponse), nil
}
