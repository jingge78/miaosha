package request

type OrderReq struct {
	Title        string  `json:"title" binding:"required"`
	RealName     string  `json:"real_name" binding:"required"`
	UserPhone    string  `json:"user_phone" binding:"required"`
	UserAddress  string  `json:"user_address" binding:"required"`
	CartId       string  `json:"cart_id" binding:"required"`
	FreightPrice float64 `json:"freight_price" binding:"required"`
	TotalNum     int64   `json:"total_num" binding:"required"`
	TotalPrice   float64 `json:"total_price" binding:"required"`
	TotalPostage float64 `json:"total_postage" binding:"required"`
	PayPrice     float64 `json:"pay_price" binding:"required"`
	PayPostage   float64 `json:"pay_postage" binding:"required"`
}

type OrderList struct {
	RealName string `json:"real_name" binding:"required"`
}
type OrderListAll struct{}
