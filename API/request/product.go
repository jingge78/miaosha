package request

type ProductDetailRequest struct {
	ProductId int `form:"productId" json:"productId" binding:"required"`
}

type EsSearchByKeyWordRequest struct {
	KeyWord string `form:"keyWord"`
}
type ProductPriceFind struct {
	Price float64 `json:"price" binding:"required"`
}
