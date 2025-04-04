package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/pkg"
	"miaosha-jjl/common/proto/user"
	"strconv"
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
	update, err := client.Update(c, &user.UpdateRequest{
		Account:  data.Account,
		Password: data.Password,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if update.Success == false {
		response.CurrencyErrorResponse(c, "失败")
		return
	}
	response.CurrencySuccessResponse(c, "用户密码修改成功", nil)

}

func SendSms(c *gin.Context) {
	var data request.SendSms
	err := c.ShouldBind(&data)
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	sms, err := client.SendSms(c, &user.SendSmsRequest{
		Mobile: data.Mobile,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if sms == nil {
		response.CurrencyErrorResponse(c, "用户密码找回失败")
		return
	}
	response.CurrencySuccessResponse(c, "验证码发送成功成功", map[string]interface{}{"user_sms": sms})
}

func PassWordRecovery(c *gin.Context) {
	var data request.PassWordRecovery
	err := c.ShouldBind(&data)
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	recovery, err := client.PasswordRecovery(c, &user.PasswordRecoveryRequest{
		Account: data.Account,
		Mobile:  data.Mobile,
		SendSms: data.SendSms,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if recovery == nil {
		response.CurrencyErrorResponse(c, "用户密码找回失败")
		return
	}
	response.CurrencySuccessResponse(c, "用户密码找回成功", map[string]interface{}{"user_recovery": recovery})
}
func UserSignIn(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	SignDate := c.Query("sign_date") //测试时间日期(如:2025-03-28)
	in, err := client.UserSignIn(c, &user.SignInRequest{
		UserId:   int32(userId),
		SignDate: SignDate,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "用户签到成功", map[string]interface{}{"user_sign in": in})
}
func UserDetail(c *gin.Context) {
	Uid := c.GetUint("userId")
	users, err := client.UserDetail(c, &user.UserDetailRequest{
		Uid: uint64(Uid),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	if users == nil {
		response.CurrencyErrorResponse(c, "用户详情展示失败")
		return
	}
	response.CurrencySuccessResponse(c, "用户详情展示成功", map[string]interface{}{"users": users})
}
func MakeupSignIn(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	SignDate := c.Query("sign_date") //测试时间日期(如:2025-03-28)
	in, err := client.MakeupSignIn(c, &user.MakeupSignInRequest{
		UserId:   int32(userId),
		SignDate: SignDate,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "用户补签成功", map[string]interface{}{"makeup_Sign_In": in})

}

func UserImproveInformation(c *gin.Context) {
	var data request.UserImproveInformation
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrorResponse(c, "参数错误")
		return
	}
	Uid := c.GetUint("userId")
	improve, err := client.UserImproveInformation(c, &user.UserImproveInformationRequest{
		Nickname: data.Nickname,
		Avatar:   data.Avatar,
		Birthday: data.Birthday,
		Address:  data.Address,
		Uid:      uint64(Uid),
	})
	if err != nil {
		response.CurrencyErrorResponse(c, "用户完善信息失败")
		return
	}
	response.CurrencySuccessResponse(c, "用户完善信息成功", map[string]interface{}{"user_improve_information": improve})

}
