package register

import (
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/product"
	"miaosha-jjl/common/proto/store_cart"
	"miaosha-jjl/common/proto/user"
	"miaosha-jjl/common/proto/user_enter"
	"miaosha-jjl/service/product_service"
	"miaosha-jjl/service/shipping_address_service"
	"miaosha-jjl/service/store_cart_service"
	"miaosha-jjl/service/user_enter_service"
	"miaosha-jjl/service/user_service"
)

func GrpcRegister(server *grpc.Server) {
	user_enter.RegisterUserEnterServer(server, &user_enter_service.ServerUserEnter{})
	store_cart.RegisterStoreCartServer(server, &store_cart_service.ServerStoreCart{})
	product.RegisterProductServer(server, &product_service.ServerProduct{})
	user.RegisterUserServer(server, &user_service.ServerUser{})
	shipping_address.RegisterShippingAddressServer(server, &shipping_address_service.ServerShippingAddress{})
}
