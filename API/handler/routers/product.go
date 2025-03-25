package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/product"
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
