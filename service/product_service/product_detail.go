package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

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
	}, nil
}
