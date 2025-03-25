package request

type ProductDetailRequest struct {
	ProductId int `form:"productId" json:"productId" binding:"required"`
}
type ProductPriceFind struct {
	Price float64 `json:"price" binding:"required"`
}
