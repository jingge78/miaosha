package register

import (
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/order"
	"miaosha-jjl/common/proto/product"
	"miaosha-jjl/common/proto/user"
	"miaosha-jjl/service/order_service"
	"miaosha-jjl/service/product_service"
	"miaosha-jjl/service/user_service"
)

func GrpcRegister(server *grpc.Server) {
	product.RegisterProductServer(server, &product_service.ServerProduct{})
	user.RegisterUserServer(server, &user_service.ServerUser{})
	order.RegisterOrderServer(server, &order_service.ServerOrder{})
}
