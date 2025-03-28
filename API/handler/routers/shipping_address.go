package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/shipping_address"
	"miaosha-jjl/common/utils"
)

func AddShippingAddress(c *gin.Context) {
	var data request.AddShippingAddress
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, "参数错误")
		return
	}
	number := utils.ValidatePhoneNumber(data.PhoneNumber)
	if number == false {
		response.CurrencyErrorResponse(c, "手机号格式错误")
		return
	}
	Uid := c.GetUint("userId")
	add, err := client.AddShippingAddress(c, &shipping_address.AddShippingAddressRequest{
		UserId:          uint64(Uid),
		Recipient_Name:  data.RecipientName,
		PhoneNumber:     data.PhoneNumber,
		Province:        data.Province,
		City:            data.City,
		District:        data.District,
		DetailedAddress: data.DetailedAddress,
		IsDefault:       data.IsDefault,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "商家入驻成功", map[string]interface{}{"user_enter": add})

}
