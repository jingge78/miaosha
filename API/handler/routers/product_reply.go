package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/product_reply"
)

func ProductReplyList(c *gin.Context) {
	var data request.ProductReplyList
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	list, err := client.ProductReplyList(c, &product_reply.ProductReplyListRequest{
		ProductId: data.ProductId,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "商品评论展示成功", map[string]interface{}{"product_reply": list})
}
