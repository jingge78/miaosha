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
