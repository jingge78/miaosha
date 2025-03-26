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
	}
	product := r.Group("/product")
	{
		product.GET("/detail", ProductDetail)
		product.GET("/list", ProductList)
		product.POST("/sync/es", ProductSyncEs)
		//Es搜索（jjl）
		product.GET("/search/es", EsSearchByKeyWord)
		product.Use(pkg.JWTAuth("2209A"))
		//收藏商品展示（jjl）
		product.POST("/collect/list", GetCollectProduct)
	}
}
