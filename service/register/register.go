package register

import (
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/product"
<<<<<<< HEAD
	"miaosha-jjl/common/proto/user"
=======
	"miaosha-jjl/common/proto/user/user"
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	"miaosha-jjl/service/product_service"
	"miaosha-jjl/service/user_service"
)

func GrpcRegister(server *grpc.Server) {
<<<<<<< HEAD
	product.RegisterProductServer(server, &product_service.ServerProduct{})
	user.RegisterUserServer(server, &user_service.ServerUser{})
=======
	product.RegisterProductServer(server, product_service.ServerProduct{})
	user.RegisterUserServer(server, user_service.ServerUser{})
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
}
