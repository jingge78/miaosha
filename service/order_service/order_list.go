package order_service

import (
	"context"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/order"
)

func (s *ServerOrder) OrderList(ctx context.Context, in *order.OrderListReq) (*order.OrderListResp, error) {
	var orders model.StoreOrder
	list, err := orders.OrderList(in.RealName)
	if err != nil {
		return nil, err
	}
	var old []*order.OrderInfo
	for _, show := range list {
		ss := show.CartId[0:6] + "********" + show.CartId[13:17]
		old = append(old, &order.OrderInfo{
			Title:        show.Title,
			Uid:          int64(show.Uid),
			RealName:     show.RealName,
			UserPhone:    show.UserPhone,
			UserAddress:  show.UserAddress,
			CartId:       ss,
			FreightPrice: float32(show.FreightPrice),
			TotalNum:     int64(show.TotalNum),
			TotalPrice:   float32(show.TotalPrice),
			TotalPostage: float32(show.TotalPostage),
			PayPrice:     float32(show.PayPrice),
			PayPostage:   float32(show.PayPostage),
		})
	}
	return &order.OrderListResp{List: old}, nil
}
