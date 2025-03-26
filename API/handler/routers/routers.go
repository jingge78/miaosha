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
		//Es搜索（jjl）
		product.GET("/search/es", EsSearchByKeyWord)
		product.POST("/price", PriceFind)
		//使用中间件
		product.Use(pkg.JWTAuth("2209A"))
		//收藏商品展示（jjl）
		product.POST("/collect/list", GetCollectProduct)
		//分类展示（LiBang）
		product.GET("/category", ProductCategory)
	}
}
