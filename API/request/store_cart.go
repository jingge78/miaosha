package request

type AddStoreCart struct {
	Type              string `form:"type" json:"type" binding:"required"`
	ProductId         uint64 `form:"productId" json:"productId" binding:"required"`
	ProductAttrUnique string `form:"productAttrUnique" json:"productAttrUnique" binding:"required"`
	CartNum           uint64 `form:"cartNum" json:"cartNum" binding:"required"`
}
