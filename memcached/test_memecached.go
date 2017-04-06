package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache" //添加go版本memcached客戶端庫
)

func TestGet(k string, mc *memcache.Client) {
	it, err := mc.Get(k)
	if err != nil {
		panic(err)
	}
	fmt.Println(it.Expiration, it.Flags, it.Key, string(it.Value))
}

// test Add function
func TestAdd(mc *memcache.Client) {
	it := &memcache.Item{Key: "honghong", Value: []byte("hello")}
	mc.Add(it)
	TestGet("honghong", mc)
}

func TestSet(mc *memcache.Client) {
	// test Set function
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	TestGet("foo", mc)
}

func TestIncrement(mc *memcache.Client) {
	mc.Set(&memcache.Item{Key: "number", Value: []byte("1")})
	mc.Increment("number", 1)
	TestGet("number", mc)
}

func TestDecrement(mc *memcache.Client) {
	//mc.Set(&memcache.Item{Key: "number", Value: []byte("1")})
	mc.Decrement("number", 1)
	TestGet("number", mc)
}

func main() {
	// 创建客户端, 连接到服务器
	mc := memcache.New("127.0.0.1:11211") //

	TestAdd(mc)
	TestSet(mc)
	TestIncrement(mc) //递增
	TestDecrement(mc) //递减
	TestDecrement(mc)
	TestDecrement(mc)
	TestDecrement(mc)

	mc.DeleteAll() //全部删除
	mc.FlushAll()

}
