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
type WebsiteProductReq struct {
	CategoryId int `form:"category_id" json:"category_id" `
	Page       int `form:"page" json:"page" binding:"required"`
	PageSize   int `form:"page_size" json:"page_size" binding:"required"`
}
