package shipping_address_service

import (
	"context"
	"errors"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/shipping_address"
)

func (a *ServerShippingAddress) AddShippingAddress(ctx context.Context, req *shipping_address.AddShippingAddressRequest) (*shipping_address.AddShippingAddressResponse, error) {
	shippingAddress := model.ShippingAddress{
		UserId:          req.UserId,
		RecipientName:   req.Recipient_Name,
		PhoneNumber:     req.PhoneNumber,
		Province:        req.Province,
		City:            req.City,
		District:        req.District,
		DetailedAddress: req.DetailedAddress,
		IsDefault:       req.IsDefault,
	}
	err := shippingAddress.AddShippingAddress()
	if err != nil {
		return nil, err
	}
	if shippingAddress.Id == 0 {
		return nil, errors.New("添加地址失败")
	}
	return &shipping_address.AddShippingAddressResponse{ShippingAddressId: shippingAddress.Id}, nil
}
