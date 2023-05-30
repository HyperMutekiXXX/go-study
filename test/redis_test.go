package test

import (
	"fmt"
	"github.com/go-study/dao"
	"testing"
)

func TestRedisString(t *testing.T) {
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
	dao.SetKeyValue("newkey", "fuck you", 0)
	value := dao.GetByKey("newkey")
	fmt.Println("newkey", value)
	dao.DelByKey("newkey")
}

func TestHash(t *testing.T) {
	dao.SetHash("nnn", "asa", 1321)
	dao.SetHash("nnn", "dsdfa", "weffwe")
	dao.SetHash("mm", "fs", 421)
	data := dao.GetHash("nnn", "asa")
	fmt.Println(data)
	dataAll := dao.GetAllHash("nnn")
	fmt.Println(dataAll)
}
