package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

func (p *ServerProduct) GetAllProduct(ctx context.Context, req *product.GetAllProductRequest) (*product.GetAllProductResponse, error) {
	var products model.Product
	allProduct, err := products.GetAllProduct()
	if err != nil {
		return nil, err
	}
	if allProduct == nil {
		return nil, errors.New("未查询到商品")
	}
	var allProductList []*product.ProductDetailResponse
	for _, productOne := range allProduct {
		allProductList = append(allProductList, &product.ProductDetailResponse{
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
	return &product.GetAllProductResponse{AllProductResponse: allProductList}, nil
}
