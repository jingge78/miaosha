package request

type AddShippingAddress struct {
	RecipientName   string `form:"recipientName" json:"recipientName" binding:"required"`
	PhoneNumber     string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	Province        string `form:"province" json:"province" binding:"required"`
	City            string `form:"city" json:"city" binding:"required"`
	District        string `form:"district" json:"district" binding:"required"`
	DetailedAddress string `form:"detailedAddress" json:"detailedAddress" binding:"required"`
	IsDefault       uint64 `form:"isDefault" json:"isDefault" binding:"required"`
}
