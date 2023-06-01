package test

import (
	"context"
	"fmt"
	"github.com/go-study/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := service.NewProdServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	resp, err := productServiceClient.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 233})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
}

func TestService(t *testing.T) {
	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, new(service.ProdService))
	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatalln("服务监听端口失败", err)
	}
	server.Serve(listen)
}
