package request

type AddUserEnter struct {
	Province     string `form:"province" json:"province" binding:"required"`
	City         string `form:"city" json:"city" binding:"required"`
	District     string `form:"district" json:"district" binding:"required"`
	Address      string `form:"address" json:"address" binding:"required"`
	MerchantName string `form:"merchantName" json:"merchantName" binding:"required"`
	LinkTel      string `form:"linkTel" json:"linkTel" binding:"required"`
	Charter      string `form:"charter" json:"charter" binding:"required"`
}
