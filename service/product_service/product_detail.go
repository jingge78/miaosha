package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

type ServerProduct struct {
	product.UnimplementedProductServer
}

func (p ServerProduct) ProductDetail(ctx context.Context, in *product.ProductDetailReq) (*product.ProductDetailResp, error) {
	address, err := ProductDetail(in)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func ProductDetail(in *product.ProductDetailReq) (*product.ProductDetailResp, error) {
	p := model.Product{}
	err := p.ProductDetailReq(in.ProductId)
	if err != nil {
		return nil, errors.New("查询商品失败")
	}
	if p.Id == 0 {
		return nil, errors.New("商品详情展示失败")
	}
	return &product.ProductDetailResp{
		MerId:     p.MerId,
		Image:     p.Image,
		StoreName: p.StoreName,
		StoreInfo: p.StoreInfo,
		Price:     float32(p.Price),
		VipPrice:  float32(p.VipPrice),
		Postage:   float32(p.IsPostage),
		Stock:     p.Stock,
		IsShow:    p.IsShow,
		IsPostage: p.IsPostage,
		IsSeckill: p.IsSeckill,
		Browse:    p.Browse,
	}, nil
}
