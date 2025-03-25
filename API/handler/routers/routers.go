package routers
<<<<<<< HEAD

import "github.com/gin-gonic/gin"

func LoadRouters(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", UserLogin)
		user.POST("/register", UserRegister)
		user.POST("/update", Update)
	}
	product := r.Group("/product")
	{
		product.GET("/detail", ProductDetail)
		product.GET("/list", ProductList)
		product.POST("/sync/es", ProductSyncEs)
	}
}
=======
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
