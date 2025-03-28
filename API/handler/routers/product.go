package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/product"
	"strconv"
)

func ProductDetail(c *gin.Context) {
	var data request.ProductDetailRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	detail, err := client.ProductDetail(c, &product.ProductDetailRequest{ProductId: uint64(data.ProductId)})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "查询成功", map[string]interface{}{"product_detail": detail})
}
func ProductList(c *gin.Context) {
	allProduct, err := client.GetAllProduct(c, &product.GetAllProductRequest{})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if allProduct.AllProductResponse == nil {
		response.CurrencyErrorResponse(c, "暂未查询到商品")
		return
	}
	response.CurrencySuccessResponse(c, "查询成功", map[string]interface{}{"product_list": allProduct.AllProductResponse})
}
func ProductSyncEs(c *gin.Context) {
	res, err := client.EsAddProduct(c, &product.EsAddProductRequest{})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if !res.Success {
		response.CurrencyErrorResponse(c, "同步失败")
		return
	}
	response.CurrencySuccessResponse(c, "同步成功", nil)
}

func EsSearchByKeyWord(c *gin.Context) {
	var data request.EsSearchByKeyWordRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	res, err := client.EsSearchByKeyWord(c, &product.EsSearchByKeyWordRequest{
		KeyWord: data.KeyWord,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "查询成功", map[string]interface{}{"res": res.EsSearchByKeyWordResponse})
}
func PriceFind(c *gin.Context) {
	var data request.ProductPriceFind
	err := c.ShouldBind(&data)
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	find, err := client.PriceFind(c, &product.PriceFindRequest{
		Price: float32(data.Price),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, "查找失败")
		return
	}
	response.CurrencySuccessResponse(c, "成功", map[string]interface{}{"price_find": find.List})
}

func GetCollectProduct(c *gin.Context) {
	userId := c.GetUint("userId")
	collectProduct, err := client.GetCollectProduct(c, &product.GetCollectProductRequest{
		UserId: int64(userId),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "查询成功", map[string]interface{}{"collectProduct": collectProduct.GetCollectProductResponse})
}

func ProductCategory(c *gin.Context) {
	parentId, _ := strconv.Atoi(c.Query("parent_id"))
	find, err := client.GetProductCategory(c, &product.ProductCategoryRequest{
		ParentId: int32(parentId),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, "展示失败")
		return
	}
	response.CurrencySuccessResponse(c, "分类展示成功", map[string]interface{}{"product_category": find})

}
func WebsiteProductList(c *gin.Context) {
	var data request.WebsiteProductReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.CurrencyErrorResponse(c, "参数请求出错")
		return
	}
	list, err := client.WebsiteProductList(c, &product.WebsiteProductListRequest{
		CategoryId: int32(data.CategoryId),
		Page:       int32(data.Page),
		PageSize:   int32(data.PageSize),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, "展示失败")
		return
	}
	response.CurrencySuccessResponse(c, "分类展示成功", map[string]interface{}{"website_product_list": list})
}
func ProductSort(c *gin.Context) {

	//IsShow是商品状态  0下架 1上架
	IsShow, _ := strconv.Atoi(c.Query("is_show"))
	list, err := client.ProductSort(c, &product.ProductSortRequest{
		IsShow: int64(IsShow),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, "展示失败")
		return
	}
	response.CurrencySuccessResponse(c, "分类展示成功", map[string]interface{}{"product_sort": list})
}
func GroupByProductList(c *gin.Context) {
	list, err := client.GroupByProductList(c, &product.GroupByProductListRequest{})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "团购商品查询成功", map[string]interface{}{"group_list": list.GroupByProductListResponse})
}

func ProductRanking(c *gin.Context) {
	_, err := client.ProductRanking(c, &product.ProductRankingRequest{})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "排行榜成功", nil)
}
