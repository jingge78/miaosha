package request

type Integration struct {
	Integral_type string `form:"integral_type" binding:"required"` //积分的类型
	Integral      string `form:"integral" binding:"required"`      //积分
	Bak           string `form:"bak" binding:"required"`           //积分补充文案
}
