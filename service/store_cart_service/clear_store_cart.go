package store_cart_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/store_cart"
)

func (s ServerStoreCart) ClearCart(ctx context.Context, in *store_cart.ClearCartRequest) (*store_cart.ClearCartResponse, error) {
	carts := &model.StoreCart{}
	err := carts.ClearCart(in.Uid)
	if err != nil {
		return nil, errors.New("清空购物车失败")
	}
	return &store_cart.ClearCartResponse{Success: true}, nil
}
