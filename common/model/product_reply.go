package model

import "miaosha-jjl/common/global"

type ProductReply struct {
	Id                   int32  `gorm:"column:id;type:int;comment:评论ID;primaryKey;not null;" json:"id"`                                              // 评论ID
	Uid                  int32  `gorm:"column:uid;type:int;comment:用户ID;not null;" json:"uid"`                                                       // 用户ID
	Oid                  int32  `gorm:"column:oid;type:int;comment:订单ID;not null;" json:"oid"`                                                       // 订单ID
	Unique               int32  `gorm:"column:unique;type:int;comment:父级id;not null;" json:"unique"`                                                 // 父级id
	ProductId            int32  `gorm:"column:product_id;type:int;comment:商品id;not null;" json:"product_id"`                                         // 商品id
	ReplyType            string `gorm:"column:reply_type;type:varchar(32);comment:某种商品类型(普通商品、秒杀商品）;not null;default:product;" json:"reply_type"`    // 某种商品类型(普通商品、秒杀商品）
	ProductScore         int8   `gorm:"column:product_score;type:tinyint(1);comment:商品分数;not null;" json:"product_score"`                            // 商品分数
	ServiceScore         int8   `gorm:"column:service_score;type:tinyint(1);comment:服务分数;not null;" json:"service_score"`                            // 服务分数
	Comment              string `gorm:"column:comment;type:varchar(512);comment:评论内容;not null;" json:"comment"`                                      // 评论内容
	Pics                 string `gorm:"column:pics;type:text;comment:评论图片;not null;" json:"pics"`                                                    // 评论图片
	AddTime              int32  `gorm:"column:add_time;type:int;comment:评论时间;not null;" json:"add_time"`                                             // 评论时间
	MerchantReplyContent string `gorm:"column:merchant_reply_content;type:varchar(300);comment:管理员回复内容;default:NULL;" json:"merchant_reply_content"` // 管理员回复内容
	MerchantReplyTime    int32  `gorm:"column:merchant_reply_time;type:int;comment:管理员回复时间;default:NULL;" json:"merchant_reply_time"`                // 管理员回复时间
	IsDel                uint8  `gorm:"column:is_del;type:tinyint UNSIGNED;comment:0未删除1已删除;not null;default:0;" json:"is_del"`                      // 0未删除1已删除
	IsReply              int8   `gorm:"column:is_reply;type:tinyint(1);comment:0未回复1已回复;not null;default:0;" json:"is_reply"`                        // 0未回复1已回复
	Nickname             string `gorm:"column:nickname;type:varchar(64);comment:用户名称;not null;" json:"nickname"`                                     // 用户名称
	Avatar               string `gorm:"column:avatar;type:varchar(255);comment:用户头像;not null;" json:"avatar"`                                        // 用户头像
}

func (r *ProductReply) GetProductReplyList() (list []*ProductReply, err error) {
	err = global.DB.Debug().Table("product_reply").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
