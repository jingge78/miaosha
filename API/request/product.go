package request

type ProductDetailRequest struct {
	ProductId int `form:"productId" json:"productId" binding:"required"`
}
<<<<<<< HEAD
type EsSearchByKeyWordRequest struct {
	KeyWord string `form:"keyWord"`
=======
type ProductPriceFind struct {
	Price float64 `json:"price" binding:"required"`
>>>>>>> 191b45c85ebaaf39781b1967de285d914c6581a1
}
