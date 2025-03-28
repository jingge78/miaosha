package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
	"strconv"
)

func (p *ServerProduct) GroupByProductList(ctx context.Context, req *product.GroupByProductListRequest) (*product.GroupByProductListResponse, error) {
	var groupProduct model.StoreCombination
	list, err := groupProduct.GroupByProductList()
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("暂未查询到团购商品")
	}
	var groupProductList []*product.GroupByProduct
	for _, item := range list {
		groupProductList = append(groupProductList, &product.GroupByProduct{
			ProductId: int64(item.ProductId),
			Image:     item.Image,
			Images:    item.Images,
			Title:     item.Title,
			People:    int64(item.People),
			Info:      item.Info,
			Price:     float32(item.Price),
			Stock:     int64(item.Stock),
			IsPostage: int64(item.IsPostage),
			Postage:   float32(item.Postage),
			StartTime: strconv.Itoa(int(item.StartTime)),
			StopTime:  strconv.Itoa(int(item.StopTime)),
			Quota:     int64(item.Quota),
		})
	}
	return &product.GroupByProductListResponse{GroupByProductListResponse: groupProductList}, nil
}
