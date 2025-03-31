package request

type AddCoupon struct {
	Title       string  `json:"title" form:"title" binding:"required"`
	Integral    uint64  `json:"integral" form:"integral" binding:"required"`
	CouponPrice float64 `json:"couponPrice" form:"couponPrice" binding:"required"`
	UseMinPrice float64 `json:"useMinPrice" form:"useMinPrice" binding:"required"`
	CouponTime  uint64  `json:"couponTime" form:"couponTime" binding:"required"`
	ProductId   string  `json:"productId" form:"productId" binding:"required"`
	CategoryId  uint64  `json:"categoryId" form:"categoryId" binding:"required"`
	Type        uint64  `json:"type" form:"type" binding:"required"`
}
