package main

import (
	"github.com/go-study/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, new(service.ProdService))
	listen, err := net.Listen("tcp", "8082")
	if err != nil {
		log.Fatalln("服务监听端口失败", err)
	}
	server.Serve(listen)
}
