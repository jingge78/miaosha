package request

type StaffService struct {
	Uid          int32  `form:"uid" binding:"required"`
	Avatar       string `form:"avatar" binding:"required"`
	StaffName    int32  `form:"StaffName" binding:"required"`
	Phone        string `form:"phone" binding:"required"`
	VerifyStatus int32  `form:"verifystatus" binding:"required"`
	status       int32  `form:"status" binding:"status"`
	Addtime      int32  `form:"addtime" binding:"add_time"`
	StoreId      int32  `form:"storeid" binding:"storeid"`
	Status       int32  `form:"status" binding:"status"`
}
