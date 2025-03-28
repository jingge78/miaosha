package order_service

import (
	"context"
	"github.com/google/uuid"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/pkg"
	"miaosha-jjl/common/proto/order"
	"strconv"
)

type ServerOrder struct {
	order.UnimplementedOrderServer
}

func (s *ServerOrder) CreateOrder(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	uid := uuid.New().String()
	m := model.StoreOrder{
		OrderId:      uid,
		Title:        in.Title,
		Uid:          uint32(in.Uid),
		RealName:     in.RealName,
		UserPhone:    in.UserPhone,
		UserAddress:  in.UserAddress,
		CartId:       in.CartId,
		FreightPrice: float64(in.FreightPrice),
		TotalNum:     uint32(in.TotalNum),
		TotalPrice:   float64(in.TotalPrice),
		TotalPostage: float64(in.TotalPostage),
		PayPrice:     float64(in.PayPrice),
		PayPostage:   float64(in.PayPostage),
	}
	err := m.CreateOrder()
	if err != nil {
		return nil, err
	}
	pay := pkg.NewPay()
	price := strconv.FormatFloat(float64(in.PayPrice), 'f', 2, 64)
	s2 := pay.Pay(in.Title, uid, price)
	return &order.CreateOrderResp{Result: s2}, err
}
