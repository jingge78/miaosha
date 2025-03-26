package store_cart_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/store_cart"
)

type ServerStoreCart struct {
	store_cart.UnimplementedStoreCartServer
}

func (s ServerStoreCart) AddStoreCart(ctx context.Context, in *store_cart.AddStoreCartRequest) (*store_cart.AddStoreCartResponse, error) {
	// 判断商品是否存在和下架
	p := model.Product{}
	err := p.ProductDetailReq(in.ProductId)
	if err != nil {
		return nil, err
	}
	if p.Id == 0 {
		return nil, errors.New("未查询到商品信息")
	}
	// 商品是否下架
	if p.IsShow == 0 {
		return nil, errors.New("该商品已经下架")
	}
	// 判断商品库存是否充足
	if p.Stock < in.CartNum {
		return nil, errors.New("商品库存不足")
	}

	// 判断购物车是否存在该商品，如果存在，则数量累加一，如果不在，则添加购物车
	c := model.StoreCart{}
	err = c.GetStoreCart(in.ProductId)
	if err != nil {
		return nil, err
	}
	if c.Id == 0 {
		c = model.StoreCart{
			Uid:               uint32(in.Uid),
			Type:              in.Type,
			ProductId:         uint32(in.ProductId),
			ProductAttrUnique: in.ProductAttrUnique,
			CartNum:           uint16(in.CartNum),
		}
		err = c.AddStoreCart()
		if err != nil {
			return nil, err
		}
		if c.Id == 0 {
			return nil, errors.New("添加购物车失败")
		}
	} else {
		err = c.UpdateStoreCartProductNum(in.ProductId, in.CartNum)
		if err != nil {
			return nil, errors.New("购物车数量累加失败")
		}
	}
	return &store_cart.AddStoreCartResponse{StoreCartId: c.Id}, nil
}
