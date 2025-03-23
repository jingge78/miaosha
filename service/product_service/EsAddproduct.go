package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/pkg"
	"miaosha-jjl/common/proto/product"
)

func (p *ServerProduct) EsAddProduct(ctx context.Context, req *product.EsAddProductRequest) (*product.EsAddProductResponse, error) {
	var products model.Product
	allProduct, err := products.GetAllProduct()
	if err != nil {
		return nil, err
	}
	if allProduct == nil {
		return nil, errors.New("未查询到商品 无法同步")
	}
	var productSetEs []*product.ProductDetailResponse
	for _, productOne := range allProduct {
		productSetEs = append(productSetEs, &product.ProductDetailResponse{
			Id:        int64(productOne.Id),
			Image:     productOne.Image,
			StoreName: productOne.StoreName,
			StoreInfo: productOne.StoreInfo,
			Price:     float32(productOne.Price),
			Postage:   float32(productOne.Postage),
			IsPostage: int64(productOne.IsPostage),
			Stock:     int64(productOne.Stock),
			Browse:    int64(productOne.Browse),
		})
	}
	err = pkg.CreateEs(productSetEs)
	if err != nil {
		return nil, err
	}
	return &product.EsAddProductResponse{Success: true}, nil
}
