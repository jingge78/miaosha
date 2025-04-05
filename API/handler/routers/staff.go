package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/API/client"
	"miaosha-jjl/API/request"
	"miaosha-jjl/API/response"
	"miaosha-jjl/common/proto/staff"
)

func Staffs(c *gin.Context) {
	var data request.StaffService
	err := c.ShouldBind(&data)
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	uid := c.GetUint("uid")
	staffClient, err := client.StaffEnd(c, &staff.StaffRequest{
		Uid:          int32(int64(uid)),
		Avatar:       data.Avatar,
		StoreId:      data.StoreId,
		StaffName:    data.StaffName,
		Phone:        data.Phone,
		VerifyStatus: data.VerifyStatus,
		Status:       data.Status,
		AddTime:      data.Addtime,
	})
	if err != nil {
		response.CurrencyErrorResponse(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"Msg":  "添加店员成功",
		"Code": 200,
		"Data": staffClient,
	})

}
