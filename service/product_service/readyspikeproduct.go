package product_service

import (
	"context"
	"fmt"
	"github.com/go-errors/errors"
	"miaosha-jjl/common/global"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
	"time"
)

func (p *ServerProduct) AddSpikeProduct(ctx context.Context, in *product.AddSpikeProductReq) (*product.AddSpikeProductResp, error) {
	var products model.Product
	findProduct, err := products.FindProduct(int(in.ProductId))
	if err != nil {
		return nil, errors.New("查询商品失败")
	}
	if findProduct.Id == 0 {
		return nil, errors.New("商品不存在")
	}

	if int64(findProduct.Stock)-in.SpikeNum <= 0 {
		return nil, errors.New("商品库存不足")
	}

	Num := int64(findProduct.Stock) - in.SpikeNum

	tx := global.DB.Begin()

	err = products.UpdateProductNum(int(in.ProductId), int(Num))
	if err != nil {
		tx.Rollback()
		return nil, errors.New("修改库存失败")
	}
	m := model.SpikeProduct{
		ProductId:    int32(findProduct.Id),
		ProductName:  findProduct.StoreName,
		ProductPrice: findProduct.Price,
		SpikePrice:   float64(in.SpikePrice),
		SpikeNum:     int32(in.SpikeNum),
		StartTime:    in.StartTime,
		EndTime:      in.EndTime,
	}
	err = m.AddSpikeProduct()
	if err != nil {
		tx.Rollback()
		return nil, errors.New("商品预热失败")
	}

	key := fmt.Sprintf("SpikeProduct:%d", m.ID)
	lockKey := fmt.Sprintf("Lock:SpikeProduct:%d", m.ID)
	lockValue := "locked"
	_, err = global.Rdb.SetNX(ctx, lockKey, lockValue, time.Minute*60).Result()
	if err != nil {
		return nil, errors.New("获取锁失败")
	}

	defer func() {
		err = global.Rdb.Del(ctx, lockKey).Err()
		if err != nil {
			fmt.Println("释放锁失败")
			return
		}
	}()

	err = global.Rdb.Del(ctx, key).Err()
	if err != nil {
		return nil, errors.New("删除库存失败")
	}

	for i := 0; i < int(in.SpikeNum); i++ {
		err = global.Rdb.LPush(ctx, key, m.ID).Err()
		if err != nil {
			return nil, errors.New("添加库存失败")
		}
	}

	tx.Commit()
	return &product.AddSpikeProductResp{Success: "商品预热成功"}, err
}
