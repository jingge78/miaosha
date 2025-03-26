package product_service

import (
	"context"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

func (p *ServerProduct) ProductSort(ctx context.Context, in *product.ProductSortRequest) (*product.ProductSortResponse, error) {

	//商品可以根据销量，价格，评价，和最新上架来进行商品排序

	var products model.Product
	price, err := products.ProductSortByIsShowOrPrice(int(in.IsShow))
	if err != nil {
		return nil, err
	}
	var Price []*product.ProductMsg
	for _, v := range price {
		Price = append(Price, &product.ProductMsg{
			Id:        int64(v.Id),
			Image:     v.Image,
			StoreName: v.StoreName,
			StoreInfo: v.StoreInfo,
			Price:     float32(v.Price),
			Postage:   float32(v.Postage),
			IsPostage: int64(v.IsPostage),
			Stock:     int64(v.Stock),
			Browse:    int64(v.Browse),
		})
	}

	return &product.ProductSortResponse{
		Products: Price,
	}, nil
}
