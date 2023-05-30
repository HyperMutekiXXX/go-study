package test

import (
	"fmt"
	"github.com/go-study/service"
	"google.golang.org/protobuf/proto"

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
