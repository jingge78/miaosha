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
		user.POST("/send", SendSms)
		user.POST("/recovery", PassWordRecovery)
		user.Use(pkg.JWTAuth("2209AGroup3"))
		user.POST("/sign", UserSignIn)
		user.POST("/makeupSignIn", MakeupSignIn)
		user.POST("/update", Update)
		user.PUT("/improve", UserImproveInformation)
		user.GET("/detail", UserDetail)
	}
	order := r.Group("/order")
	{
		order.Use(pkg.JWTAuth("2209AGroup3"))
		order.POST("/orderCreate", OrderCreate)
		order.POST("/orderList", OrderList)
		order.POST("/orderListAll", OrderListAll)
	}
	product := r.Group("/product")
	{
		product.GET("/detail", ProductDetail)
		product.GET("/list", ProductList)
		product.POST("/sync/es", ProductSyncEs)
		product.POST("/productRanking", ProductRanking)
		product.POST("/spikeProduct", SpikeProduct)
		//Es搜索（jjl）
		product.GET("/search/es", EsSearchByKeyWord)
		product.POST("/price", PriceFind)

		product.GET("/website/product", WebsiteProductList) //网站商品展示
		product.GET("/sort", ProductSort)                   //网站商品排序
		//分类展示（LiBang）
		product.GET("/category", ProductCategory)
		product.GET("/filter", ProductFilter)
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

	//	r.POST("/signup/signup", Send)

	shipess := r.Group("/signup")
	{
		shipess.POST("/signup", Send)
	}
	initialition := r.Group("/initialition")
	{
		initialition.POST("/initialition", IntegralogComment)
	}

	productReply := r.Group("/product_reply")
	{
		productReply.GET("/list", ProductReplyList)
	}

	coupon := r.Group("/coupon")
	{
		coupon.POST("/add", AddCoupon)
		coupon.POST("/grant", GrantCouponUser)
	}

	staff := r.Group("/staff")
	{
		staff.POST("/staff", Staffs)
	}

}
