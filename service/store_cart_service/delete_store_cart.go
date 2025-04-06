package store_cart_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/store_cart"
)

func (s ServerStoreCart) DeleteCart(ctx context.Context, in *store_cart.DeleteCartRequest) (*store_cart.DeleteCartResponse, error) {
	carts := &model.StoreCart{}
	err := carts.DeleteCart(in.Uid, in.ProductId)
	if err != nil {
		return nil, errors.New("删除购物车商品失败")
	}
	return &store_cart.DeleteCartResponse{Success: true}, nil
}
