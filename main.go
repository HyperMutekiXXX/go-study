package main

import (
	"fmt"
	"github.com/go-study/dao"
)

func main() {
	//ctx := context.Background()
	//err := db.Rdb.Set(ctx, "new", "kkk", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//result, err := db.Rdb.Get(ctx, "new").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("new", result)
	//val, err := db.Rdb.Do(ctx, "get", "new").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("new", val)
	dao.SetKeyValue("newkey", "fuck you", 1000*1000*1000*10)
	value := dao.GetByKey("newkey")
	fmt.Println("newkey", value)
}
