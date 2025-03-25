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
