package routers

import (
	"github.com/gin-gonic/gin"
	"miaosha-jjl/common/pkg"
)

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

		product.GET("/website/product", WebsiteProductList) //网站商品展示
		product.GET("/sort", ProductSort)                   //网站商品排序
		product.Use(pkg.JWTAuth("2209A"))
		//收藏商品展示
		product.POST("/collect/list", GetCollectProduct)

		product.GET("/category", ProductCategory)

	}
}
