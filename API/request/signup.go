package request

type Login struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Register struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type PasswordRecovery struct {
	Account string `form:"account" binding:"required"`
	Mobile  string `form:"mobile" binding:"required"`
	SendSms string `form:"sendSms" binding:"required"`
}

type AddEmail struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
