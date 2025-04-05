package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/signup"
)

func Send(c *gin.Context) {
	var data request.AddEmail
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	fmt.Printf("111111111111111111111")
	add, err := client.Send(c, &signup.SignupRequest{
		Email:    data.Email,
		Password: data.Password,
	})
	fmt.Printf("2222222222222")
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	fmt.Printf("3333333333333")
	response.CurrencySuccessResponse(c, "邮箱注册成功", map[string]interface{}{"signup": add})

}
