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
		//Es搜索（jjl）
		product.GET("/search/es", EsSearchByKeyWord)
		product.POST("/price", PriceFind)

		product.GET("/website/product", WebsiteProductList) //网站商品展示
		product.GET("/sort", ProductSort)                   //网站商品排序
		//分类展示（LiBang）
		product.GET("/category", ProductCategory)
		product.GET("group/list", GroupByProductList)
		//使用中间件
		product.Use(pkg.JWTAuth("2209AGroup3"))
		//收藏商品展示（jjl）
		product.GET("/collect/list", GetCollectProduct)

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

	shippingAddress := r.Group("/shipping_address")
	{
		shippingAddress.Use(pkg.JWTAuth("2209AGroup3"))
		shippingAddress.POST("/add", AddShippingAddress)
	}

	productReply := r.Group("/product_reply")
	{
		productReply.GET("/list", ProductReplyList)
	}
}
