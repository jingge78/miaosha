package product_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/pkg"
	"miaosha-jjl/common/proto/product"
)

func (p *ServerProduct) EsSearchByKeyWord(ctx context.Context, req *product.EsSearchByKeyWordRequest) (*product.EsSearchByKeyWordResponse, error) {
	res, err := pkg.EsSearchByKeyWord(req.KeyWord)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("没有查询到内容")
	}
	return &product.EsSearchByKeyWordResponse{
		EsSearchByKeyWordResponse: res,
	}, nil
}
