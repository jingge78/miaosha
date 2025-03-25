package routers

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
<<<<<<< HEAD
		product.GET("/search/es", EsSearchByKeyWord)
=======
		product.POST("/price", PriceFind)
>>>>>>> 191b45c85ebaaf39781b1967de285d914c6581a1
	}
}
