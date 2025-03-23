package GRPC

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"strconv"
)

func RegisterGrpc(host string, port int, call func(server *grpc.Server)) {
	listen, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	call(server)
	log.Printf("grpc listening on %v", listen.Addr())
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
