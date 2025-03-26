package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

func (p *ServerProduct) GetCollectProduct(ctx context.Context, req *product.GetCollectProductRequest) (*product.GetCollectProductResponse, error) {
	var collectProduct model.EbStoreProductRelation
	getCollectProduct, err := collectProduct.GetCollectProduct(int(req.UserId))
	if err != nil {
		return nil, err
	}
	if len(getCollectProduct) == 0 {
		return nil, errors.New("您没有收藏商品")
	}
	var collectProductList []*product.CollectProduct
	for _, collectProductOne := range getCollectProduct {
		collectProductList = append(collectProductList, &product.CollectProduct{
			ProductId: int64(collectProductOne.ProductId),
			Type:      collectProduct.Type,
			Category:  collectProductOne.Category,
			AddTime:   collectProduct.AddTime.Format("2006-01-02"),
		})
	}
	return &product.GetCollectProductResponse{
		GetCollectProductResponse: collectProductList,
	}, nil
}
