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
	order := r.Group("/order")
	{
		order.Use(pkg.JWTAuth("2209AGroup3"))
		order.POST("/orderCreate", OrderCreate)
	}
	product := r.Group("/product")
	{
		product.GET("/detail", ProductDetail)
		product.GET("/list", ProductList)
		product.POST("/sync/es", ProductSyncEs)
		product.GET("/search/es", EsSearchByKeyWord)
		product.POST("/price", PriceFind)
		product.Use(pkg.JWTAuth("2209A"))
		//收藏商品展示
		product.POST("/collect/list", GetCollectProduct)
		product.GET("/category", ProductCategory)

	}
}
