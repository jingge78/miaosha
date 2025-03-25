package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

<<<<<<< HEAD
type ServerProduct struct {
	product.UnimplementedProductServer
}

<<<<<<< HEAD
=======
>>>>>>> f7cd67524df94e75b5a629c18176469208cedc32
func (p *ServerProduct) ProductDetail(ctx context.Context, in *product.ProductDetailRequest) (*product.ProductDetailResponse, error) {
	productDetail := model.Product{}
	err := productDetail.ProductDetailReq(in.ProductId)
	if err != nil {
		return nil, err
	}
	if productDetail.Id == 0 {
		return nil, errors.New("未查询到商品信息")
	}
	return &product.ProductDetailResponse{
		Id:        int64(productDetail.Id),
		Image:     productDetail.Image,
		StoreName: productDetail.StoreName,
		StoreInfo: productDetail.StoreInfo,
		Price:     float32(productDetail.Price),
		Postage:   float32(productDetail.Postage),
		Stock:     int64(productDetail.Stock),
		IsPostage: int64(productDetail.IsPostage),
		Browse:    int64(productDetail.Browse),
=======
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
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	}, nil
}
