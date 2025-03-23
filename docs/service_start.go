package main

import (
	"google.golang.org/grpc"
	"miaosha-jjl/common/GRPC"
	"miaosha-jjl/common/initialition"
	"miaosha-jjl/service/register"
)

func main() {
	initialition.Init()
	GRPC.RegisterGrpc("127.0.0.1", 8081, func(server *grpc.Server) {
		register.GrpcRegister(server)
	})
}
