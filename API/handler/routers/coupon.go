package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/coupon"
)

func AddCoupon(c *gin.Context) {
	var data request.AddCoupon
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	coupons, err := client.AddCoupon(c, &coupon.AddCouponRequest{
		Title:       data.Title,
		Integral:    data.Integral,
		CouponPrice: float32(data.CouponPrice),
		UseMinPrice: float32(data.UseMinPrice),
		CouponTime:  data.CouponTime,
		ProductId:   data.ProductId,
		CategoryId:  data.CategoryId,
		Type:        data.Type,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if coupons.CouponId == 0 {
		response.CurrencyErrorResponse(c, "平台优惠券添加失败")
		return
	}
	response.CurrencySuccessResponse(c, "平台优惠券添加成功", map[string]interface{}{"coupon": coupons})
}

func GrantCouponUser(c *gin.Context) {
	var data request.GrantCouponUser
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	coupons, err := client.GrantCouponUser(c, &coupon.GrantCouponUserRequest{
		Uid:         data.AddTime,
		CouponTitle: data.CouponTitle,
		CouponPrice: float32(data.CouponPrice),
		UseMinPrice: float32(data.UseMinPrice),
		AddTime:     data.AddTime,
		EndTime:     data.EndTime,
		Type:        data.Type,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if coupons.Success == false {
		response.CurrencyErrorResponse(c, "用户获取平台发放优惠券失败")
		return
	}
	response.CurrencySuccessResponse(c, "用户获取平台发放优惠券成功", map[string]interface{}{"coupon_user": coupons})
}
