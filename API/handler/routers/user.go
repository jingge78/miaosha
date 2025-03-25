package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/pkg"
	"miaosha-jjl/common/proto/user"
)

func UserLogin(c *gin.Context) {
	var data request.UserLoginRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	login, err := client.UserLogin(c, &user.LoginRequest{
		Account:  data.Account,
		Password: data.Password,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	token, err := pkg.NewJWT("2209AGroup3").CreateToken(pkg.CustomClaims{
		ID: uint(login.UserId),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "登录成功", map[string]interface{}{"token": token})
}
func UserRegister(c *gin.Context) {
	var dara request.UserRegisterRequest
	if err := c.ShouldBind(&dara); err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	register, err := client.UserRegister(c, &user.RegisterRequest{
		Account:  dara.Account,
		Password: dara.Password,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if !register.Success {
		response.CurrencyErrorResponse(c, "用户注册失败")
		return
	}
	response.CurrencySuccessResponse(c, "用户注册成功", nil)
}
func Update(c *gin.Context) {
	var data request.Update
	err := c.ShouldBind(&data)
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	_, err = client.Update(c, &user.UpdateRequest{
		Account:  data.Account,
		Password: data.Password,
	})
	response.CurrencySuccessResponse(c, "用户密码修改成功", nil)

}
