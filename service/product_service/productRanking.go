package product_service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"miaosha-jjl/common/global"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

func (p *ServerProduct) ProductRanking(ctx context.Context, in *product.ProductRankingRequest) (*product.ProductRankingResponse, error) {
	var products model.Product
	allProduct, err := products.GetAllProduct()
	if err != nil {
		return nil, err
	}
	var all []*product.ProductDetailResponse
	for _, p2 := range allProduct {
		all = append(all, &product.ProductDetailResponse{
			Id:    int64(p2.Id),
			Stock: int64(p2.Stock),
		})
		global.Rdb.ZAdd(ctx, "stock_rank", &redis.Z{
			Score:  float64(p2.Stock),
			Member: p2.Id,
		})
	}

	result, err := global.Rdb.ZRevRangeWithScores(ctx, "stock_rank", 0, 4).Result()
	if err != nil {
		return nil, err
	}
	for _, z := range result {
		productId := z.Member
		stock := z.Score
		fmt.Sprintf("商品id:%d:库存:%d", productId, stock)
	}
	return nil, err
}
