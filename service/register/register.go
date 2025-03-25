package register

import (
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/order"
	"miaosha-jjl/common/proto/product"
<<<<<<< HEAD
	"miaosha-jjl/common/proto/user"
<<<<<<< HEAD
=======
	"miaosha-jjl/common/proto/user/user"
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
=======
	"miaosha-jjl/common/proto/user_enter"
	"miaosha-jjl/service/order_service"
>>>>>>> f7cd67524df94e75b5a629c18176469208cedc32
	"miaosha-jjl/service/product_service"
	"miaosha-jjl/service/user_enter_service"
	"miaosha-jjl/service/user_service"
)

func GrpcRegister(server *grpc.Server) {
<<<<<<< HEAD
<<<<<<< HEAD
=======
	user_enter.RegisterUserEnterServer(server, &user_enter_service.ServerUserEnter{})
>>>>>>> f7cd67524df94e75b5a629c18176469208cedc32
	product.RegisterProductServer(server, &product_service.ServerProduct{})
	order.RegisterOrderServer(server, &order_service.ServerOrder{})
	user.RegisterUserServer(server, &user_service.ServerUser{})
=======
	product.RegisterProductServer(server, product_service.ServerProduct{})
	user.RegisterUserServer(server, user_service.ServerUser{})
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
}
