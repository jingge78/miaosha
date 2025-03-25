package routers

import "github.com/gin-gonic/gin"

func LoadRouters(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", UserLogin)
		user.POST("/register", UserRegister)
		user.POST("/update", Update)
		user.POST("/send", SendSms)
		user.POST("/recovery", PassWordRecovery)
	}
	product := r.Group("/product")
	{
		product.GET("/detail", ProductDetail)
		product.GET("/list", ProductList)
		product.POST("/sync/es", ProductSyncEs)
		product.GET("/search/es", EsSearchByKeyWord)
		product.POST("/price", PriceFind)

		product.GET("/category", ProductCategory)

	}
}
