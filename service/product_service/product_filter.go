package product_service

import (
	"context"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

// 商品属性筛选
func (p *ServerProduct) ProductFilter(ctx context.Context, in *product.ProductFilterRequest) (*product.ProductFilterResponse, error) {

	products := model.Product{}
	filter, u, err := products.ProductFilter(in.MinPrice, in.MaxPrice, in.StoreName, in.IsPostage, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	var slice []*product.ProductMsg
	for _, v := range filter {
		slice = append(slice, &product.ProductMsg{
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
	return &product.ProductFilterResponse{
		ProductsList: slice,
		Pagination: &product.Pagination{
			Page:      int32(u.Page),
			PageSize:  int32(u.PageSize),
			Total:     int32(u.Total),
			TotalPage: int32(u.TotalPage),
		},
	}, nil
}
