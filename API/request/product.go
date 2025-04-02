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
type AddSpikeProductReq struct {
	ProductId    int     `json:"product_id" binding:"required"`
	ProductName  string  `json:"product_name" binding:"required"`
	ProductPrice float64 `json:"product_price" binding:"required"`
	SpikePrice   float64 `json:"spike_price" binding:"required"`
	SpikeNum     int64   `json:"spike_num" binding:"required"`
	StartTime    string  `json:"start_time" binding:"required"`
	EndTime      string  `json:"end_time" binding:"required"`
}
