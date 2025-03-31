package coupon_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/coupon"
)

type CouponServer struct {
	coupon.UnimplementedCouponServer
}

func (c CouponServer) AddCoupon(ctx context.Context, in *coupon.AddCouponRequest) (*coupon.AddCouponResponse, error) {
	coupons := model.Coupon{
		Title:       in.Title,
		Integral:    uint32(in.Integral),
		CouponPrice: float64(in.CouponPrice),
		UseMinPrice: float64(in.UseMinPrice),
		CouponTime:  uint32(in.CouponTime),
		Status:      0,
		ProductId:   in.ProductId,
		CategoryId:  int32(in.CategoryId),
		Type:        int8(in.Type),
	}
	err := coupons.AddCoupon()
	if err != nil {
		return nil, err
	}
	if coupons.Id == 0 {
		return nil, errors.New("平台优惠券添加失败")
	}
	return &coupon.AddCouponResponse{CouponId: uint64(coupons.Id)}, nil
}
