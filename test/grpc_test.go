package test

import (
	"fmt"
	"github.com/go-study/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"testing"
)

func TestProto(t *testing.T) {
	var password string = "asdasd"
	user := &service.User{
		Username: "mszlu",
		Age:      20,
		Password: &password,
		Address:  []string{"asdsa", "addas"},
	}
	//转换为protobuf
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser)
}

func TestService(t *testing.T) {
	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, new(service.ProdService))
	listen, err := net.Listen("tcp", "8082")
	if err != nil {
		log.Fatalln("服务监听端口失败", err)
	}
	server.Serve(listen)
}
