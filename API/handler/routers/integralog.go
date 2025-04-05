package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/initialition"
)

func IntegralogComment(c *gin.Context) {

	var data request.Integration

	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	add, err := client.Integration(c, &initialition.IntegrationRequest{
		IntegralType:  data.Integral_type,
		Integral:      data.Integral,
		Bak:           data.Bak,
		OperationTime: "",
		CreateTime:    "",
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "我的积分成功", map[string]interface{}{"initialition": add})
}
