package product_reply_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product_reply"
)

func (r *ServerProductReply) ProductReplyList(ctx context.Context, req *product_reply.ProductReplyListRequest) (*product_reply.ProductReplyListResponse, error) {
	products := model.Product{}
	err := products.ProductDetailReq(req.ProductId)
	if err != nil {
		return nil, err
	}
	if products.Id == 0 {
		return nil, errors.New("未查询到商品")
	}
	productReply := model.ProductReply{}
	list, err := productReply.GetProductReplyList()
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, errors.New("未查询到商品评论列表")
	}
	var productReplyList []*product_reply.ProductReplyList
	for _, reply := range list {
		productReplyList = append(productReplyList, &product_reply.ProductReplyList{
			ReplyId:              uint64(reply.Id),
			Uid:                  uint64(reply.Uid),
			Oid:                  uint64(reply.Oid),
			Unique:               uint64(reply.Unique),
			ProductId:            uint64(reply.ProductId),
			ReplyType:            reply.ReplyType,
			ProductScore:         uint64(reply.ProductScore),
			ServiceScore:         uint64(reply.ServiceScore),
			Comment:              reply.Comment,
			Pics:                 reply.Pics,
			AddTime:              string(reply.AddTime),
			MerchantReplyContent: reply.MerchantReplyContent,
			MerchantReplyTime:    string(reply.MerchantReplyTime),
			Nickname:             reply.Nickname,
			Avatar:               reply.Avatar,
		})
	}
	return &product_reply.ProductReplyListResponse{List: productReplyList}, nil
}
