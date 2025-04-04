package request

type UserLoginRequest struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type UserRegisterRequest struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Update struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SendSms struct {
	Mobile string `form:"mobile" binding:"required"`
}

type PassWordRecovery struct {
	Account string `form:"account" binding:"required"`
	Mobile  string `form:"mobile" binding:"required"`
	SendSms string `form:"sendSms" binding:"required"`
}

type UserImproveInformation struct {
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
	Birthday uint64 `json:"birthday" form:"birthday" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
}
