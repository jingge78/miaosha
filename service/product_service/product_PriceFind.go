package product_service

import (
	"context"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

// 价格查找//jj
func (p *ServerProduct) PriceFind(ctx context.Context, in *product.PriceFindRequest) (*product.PriceFindResponse, error) {
	var ProductPriceFind model.Product
	pro, err := ProductPriceFind.PriceFind(float64(in.Price))
	if err != nil {
		return nil, err
	}
	
	var Price []*product.PriceInfo
	for _, p2 := range pro {
		Price = append(Price, &product.PriceInfo{
			Id:        int64(p2.Id),
			Image:     p2.Image,
			StoreName: p2.StoreName,
			StoreInfo: p2.StoreInfo,
			Price:     float32(p2.Price),
			Postage:   float32(p2.Postage),
			Stock:     int64(p2.Stock),
			IsPostage: int64(p2.IsPostage),
			Browse:    int64(p2.Browse),
		})
	}
	return &product.PriceFindResponse{
		List: Price,
	}, nil
}
