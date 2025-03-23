package register

import (
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/product"
	"miaosha-jjl/common/proto/user"
	"miaosha-jjl/service/product_service"
	"miaosha-jjl/service/user_service"
)

func GrpcRegister(server *grpc.Server) {
	product.RegisterProductServer(server, &product_service.ServerProduct{})
	user.RegisterUserServer(server, &user_service.ServerUser{})
}
