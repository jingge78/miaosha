package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/store_cart"
)

func AddStoreCart(c *gin.Context) {
	var data request.AddStoreCart
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	Uid := c.GetUint("userId")
	add, err := client.AddStoreCart(c, &store_cart.AddStoreCartRequest{
		Uid:               uint64(Uid),
		Type:              data.Type,
		ProductId:         data.ProductId,
		ProductAttrUnique: data.ProductAttrUnique,
		CartNum:           data.CartNum,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "添加购物车成功", map[string]interface{}{"store_cart": add})
}

func DeleteCart(c *gin.Context) {
	var data request.DeleteCart
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	Uid := c.GetUint("userId")
	add, err := client.DeleteCart(c, &store_cart.DeleteCartRequest{
		Uid:       uint64(Uid),
		ProductId: data.ProductId,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, "删除购物车商品失败")
		return
	}
	response.CurrencySuccessResponse(c, "删除购物车商品成功", map[string]interface{}{"store_cart": add})
}

func ClearCart(c *gin.Context) {
	Uid := c.GetUint("userId")
	add, err := client.ClearCart(c, &store_cart.ClearCartRequest{
		Uid: uint64(Uid),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, "清空购物车失败")
		return
	}
	response.CurrencySuccessResponse(c, "清空购物车成功", map[string]interface{}{"store_cart": add})
}
