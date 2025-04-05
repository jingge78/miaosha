package coupon_service

import (
	"context"
	"miaosha-jjl/common/proto/coupon"
)

func (c CouponServer) GrantCouponUser(ctx context.Context, in *coupon.GrantCouponUserRequest) (*coupon.GrantCouponUserResponse, error) {
	// 平台发放给所有用户的优惠券

	// 用户登录会收到平台发放的优惠券

	// 然后用户在进行使用
	return &coupon.GrantCouponUserResponse{Success: true}, nil
}
