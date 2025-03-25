package model

type StoreOrder struct {
	Id                     uint32  `gorm:"column:id;type:int UNSIGNED;comment:订单ID;primaryKey;not null;" json:"id"`                                                          // 订单ID
	OrderId                string  `gorm:"column:order_id;type:varchar(32);comment:订单号;not null;" json:"order_id"`                                                           // 订单号
	Uid                    uint32  `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                                                                   // 用户id
	RealName               string  `gorm:"column:real_name;type:varchar(32);comment:用户姓名;not null;" json:"real_name"`                                                        // 用户姓名
	UserPhone              string  `gorm:"column:user_phone;type:varchar(18);comment:用户电话;not null;" json:"user_phone"`                                                      // 用户电话
	UserAddress            string  `gorm:"column:user_address;type:varchar(100);comment:详细地址;not null;" json:"user_address"`                                                 // 详细地址
	CartId                 string  `gorm:"column:cart_id;type:varchar(256);comment:购物车id;not null;default:[];" json:"cart_id"`                                               // 购物车id
	FreightPrice           float64 `gorm:"column:freight_price;type:decimal(8, 2);comment:运费金额;not null;default:0.00;" json:"freight_price"`                                 // 运费金额
	TotalNum               uint32  `gorm:"column:total_num;type:int UNSIGNED;comment:订单商品总数;not null;default:0;" json:"total_num"`                                           // 订单商品总数
	TotalPrice             uint32  `gorm:"column:total_price;type:decimal(8, 2) UNSIGNED;comment:订单总价;not null;default:0.00;" json:"total_price"`                            // 订单总价
	TotalPostage           uint32  `gorm:"column:total_postage;type:decimal(8, 2) UNSIGNED;comment:邮费;not null;default:0.00;" json:"total_postage"`                          // 邮费
	PayPrice               uint32  `gorm:"column:pay_price;type:decimal(8, 2) UNSIGNED;comment:实际支付金额;not null;default:0.00;" json:"pay_price"`                              // 实际支付金额
	PayPostage             uint32  `gorm:"column:pay_postage;type:decimal(8, 2) UNSIGNED;comment:支付邮费;not null;default:0.00;" json:"pay_postage"`                            // 支付邮费
	DeductionPrice         uint32  `gorm:"column:deduction_price;type:decimal(8, 2) UNSIGNED;comment:抵扣金额;not null;default:0.00;" json:"deduction_price"`                    // 抵扣金额
	CouponId               uint32  `gorm:"column:coupon_id;type:int UNSIGNED;comment:优惠券id;not null;default:0;" json:"coupon_id"`                                            // 优惠券id
	CouponPrice            uint32  `gorm:"column:coupon_price;type:decimal(8, 2) UNSIGNED;comment:优惠券金额;not null;default:0.00;" json:"coupon_price"`                         // 优惠券金额
	Paid                   uint8   `gorm:"column:paid;type:tinyint UNSIGNED;comment:支付状态;not null;default:0;" json:"paid"`                                                   // 支付状态
	PayTime                uint32  `gorm:"column:pay_time;type:int UNSIGNED;comment:支付时间;default:NULL;" json:"pay_time"`                                                     // 支付时间
	PayType                string  `gorm:"column:pay_type;type:varchar(32);comment:支付方式;not null;" json:"pay_type"`                                                          // 支付方式
	AddTime                uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:创建时间;not null;" json:"add_time"`                                                         // 创建时间
	Status                 int8    `gorm:"column:status;type:tinyint(1);comment:订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）;not null;default:0;" json:"status"` // 订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）
	RefundStatus           uint8   `gorm:"column:refund_status;type:tinyint UNSIGNED;comment:0 未退款 1 申请中 2 已退款;not null;default:0;" json:"refund_status"`                    // 0 未退款 1 申请中 2 已退款
	RefundReasonWapImg     string  `gorm:"column:refund_reason_wap_img;type:varchar(255);comment:退款图片;default:NULL;" json:"refund_reason_wap_img"`                           // 退款图片
	RefundReasonWapExplain string  `gorm:"column:refund_reason_wap_explain;type:varchar(255);comment:退款用户说明;default:NULL;" json:"refund_reason_wap_explain"`                 // 退款用户说明
	RefundReasonTime       uint32  `gorm:"column:refund_reason_time;type:int UNSIGNED;comment:退款时间;default:NULL;" json:"refund_reason_time"`                                 // 退款时间
	RefundReasonWap        string  `gorm:"column:refund_reason_wap;type:varchar(255);comment:前台退款原因;default:NULL;" json:"refund_reason_wap"`                                 // 前台退款原因
	RefundReason           string  `gorm:"column:refund_reason;type:varchar(255);comment:不退款的理由;default:NULL;" json:"refund_reason"`                                         // 不退款的理由
	RefundPrice            uint32  `gorm:"column:refund_price;type:decimal(8, 2) UNSIGNED;comment:退款金额;not null;default:0.00;" json:"refund_price"`                          // 退款金额
	DeliveryName           string  `gorm:"column:delivery_name;type:varchar(64);comment:快递名称/送货人姓名;default:NULL;" json:"delivery_name"`                                      // 快递名称/送货人姓名
	DeliveryType           string  `gorm:"column:delivery_type;type:varchar(32);comment:发货类型;default:NULL;" json:"delivery_type"`                                            // 发货类型
	DeliveryId             string  `gorm:"column:delivery_id;type:varchar(64);comment:快递单号/手机号;default:NULL;" json:"delivery_id"`                                            // 快递单号/手机号
	GainIntegral           uint32  `gorm:"column:gain_integral;type:decimal(8, 2) UNSIGNED;comment:消费赚取积分;not null;default:0.00;" json:"gain_integral"`                      // 消费赚取积分
	UseIntegral            uint32  `gorm:"column:use_integral;type:decimal(8, 2) UNSIGNED;comment:使用积分;not null;default:0.00;" json:"use_integral"`                          // 使用积分
	BackIntegral           uint32  `gorm:"column:back_integral;type:decimal(8, 2) UNSIGNED;comment:给用户退了多少积分;default:NULL;" json:"back_integral"`                            // 给用户退了多少积分
	Mark                   string  `gorm:"column:mark;type:varchar(512);comment:备注;not null;" json:"mark"`                                                                   // 备注
	IsDel                  uint8   `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;not null;default:0;" json:"is_del"`                                               // 是否删除
	Unique                 string  `gorm:"column:unique;type:char(32);comment:唯一id(md5加密)类似id;not null;" json:"unique"`                                                      // 唯一id(md5加密)类似id
	Remark                 string  `gorm:"column:remark;type:varchar(512);comment:管理员备注;default:NULL;" json:"remark"`                                                        // 管理员备注
	MerId                  uint32  `gorm:"column:mer_id;type:int UNSIGNED;comment:商户ID;not null;default:0;" json:"mer_id"`                                                   // 商户ID
	IsMerCheck             uint8   `gorm:"column:is_mer_check;type:tinyint UNSIGNED;not null;default:0;" json:"is_mer_check"`
	CombinationId          uint32  `gorm:"column:combination_id;type:int UNSIGNED;comment:拼团商品id0一般商品;default:0;" json:"combination_id"`            // 拼团商品id0一般商品
	PinkId                 uint32  `gorm:"column:pink_id;type:int UNSIGNED;comment:拼团id 0没有拼团;not null;default:0;" json:"pink_id"`                  // 拼团id 0没有拼团
	Cost                   uint32  `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本价;not null;" json:"cost"`                               // 成本价
	SeckillId              uint32  `gorm:"column:seckill_id;type:int UNSIGNED;comment:秒杀商品ID;not null;default:0;" json:"seckill_id"`                // 秒杀商品ID
	BargainId              uint32  `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价id;default:0;" json:"bargain_id"`                           // 砍价id
	VerifyCode             string  `gorm:"column:verify_code;type:varchar(12);comment:核销码;not null;" json:"verify_code"`                            // 核销码
	StoreId                int32   `gorm:"column:store_id;type:int;comment:门店id;not null;default:0;" json:"store_id"`                               // 门店id
	ShippingType           int8    `gorm:"column:shipping_type;type:tinyint(1);comment:配送方式 1=快递 ，2=门店自提;not null;default:1;" json:"shipping_type"` // 配送方式 1=快递 ，2=门店自提
	ClerkId                int32   `gorm:"column:clerk_id;type:int;comment:店员id;not null;default:0;" json:"clerk_id"`                               // 店员id
	IsChannel              uint8   `gorm:"column:is_channel;type:tinyint UNSIGNED;comment:支付渠道(0微信公众号1微信小程序);default:0;" json:"is_channel"`         // 支付渠道(0微信公众号1微信小程序)
	IsRemind               uint8   `gorm:"column:is_remind;type:tinyint UNSIGNED;comment:消息提醒;default:0;" json:"is_remind"`                         // 消息提醒
	IsSystemDel            int8    `gorm:"column:is_system_del;type:tinyint(1);comment:后台是否删除;default:0;" json:"is_system_del"`                     // 后台是否删除
}
