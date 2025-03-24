package response

import "github.com/gin-gonic/gin"

func CurrencySuccessResponse(c *gin.Context, msg string, data map[string]interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}
func CurrencyErrorResponse(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": -1,
		"msg":  msg,
		"data": nil,
	})
}
