package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/order"
)

func OrderCreate(c *gin.Context) {
	var data request.OrderReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	ui := c.GetUint("userId")
	orderClient, err := client.CreateOrder(c, &order.CreateOrderReq{
		Title:        data.Title,
		Uid:          int64(ui),
		RealName:     data.RealName,
		UserPhone:    data.UserPhone,
		UserAddress:  data.UserAddress,
		CartId:       data.CartId,
		FreightPrice: float32(data.FreightPrice),
		TotalNum:     data.TotalNum,
		TotalPrice:   float32(data.TotalPrice),
		TotalPostage: float32(data.TotalPostage),
		PayPrice:     float32(data.PayPrice),
		PayPostage:   float32(data.PayPostage),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"Msg":  "添加成功",
		"Code": 200,
		"Data": orderClient,
	})
}
