package request

type ProductReplyList struct {
	ProductId uint64 `form:"productId" json:"productId" binding:"required"`
}
