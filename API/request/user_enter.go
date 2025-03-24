package request

type AddUserEnterRequest struct {
	Province string `json:"province" form:"province" binding:"required"`
	City     string `json:"city" form:"city" binding:"required"`
	District string `json:"district" form:"district" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
	LinkTel  string `json:"linkTel" form:"linkTel" binding:"required"`
	Charter  string `json:"charter" form:"charter" binding:"required"`
}
