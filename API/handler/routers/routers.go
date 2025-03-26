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
<<<<<<< HEAD
		product.POST("/price", PriceFind)
<<<<<<< HEAD
=======
		product.Use(pkg.JWTAuth("2209A"))
		product.POST("/collect/list", GetCollectProduct)
>>>>>>> jjl
=======

		product.GET("/category", ProductCategory)

>>>>>>> origin/main
	}

	userEnter := r.Group("/user_enter")
	{
		userEnter.Use(pkg.JWTAuth("2209AGroup3"))
		userEnter.POST("/add", AddUserEnter)
	}

	storeCart := r.Group("/store_cart")
	{
		storeCart.Use(pkg.JWTAuth("2209AGroup3"))
		storeCart.POST("/add", AddStoreCart)
	}
}
