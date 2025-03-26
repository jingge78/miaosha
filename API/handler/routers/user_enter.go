package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/user_enter"
)

func AddUserEnter(c *gin.Context) {
	var data request.AddUserEnter
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	Uid := c.GetUint("userId")
	add, err := client.AddUserEnter(c, &user_enter.AddUserEnterRequest{
		Uid:          uint64(Uid),
		Province:     data.Province,
		City:         data.City,
		District:     data.District,
		Address:      data.Address,
		MerchantName: data.MerchantName,
		LinkTel:      data.LinkTel,
		Charter:      data.Charter,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "商家入驻成功", map[string]interface{}{"user_enter": add})
}
