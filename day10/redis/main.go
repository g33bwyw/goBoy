package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

//初始化redis数据库
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.10.10:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()

	return
}
func main() {
	err := initClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	//s,err := rdb.Set(ctx,"name", "wyw", 0).Result()
	//fmt.Println(s,err)
	//s1,err := rdb.Get(ctx, "name").Result()
	//fmt.Println(s1, err)
	s, err := rdb.SetNX(ctx, "id", 1, time.Second*1).Result()
	rdb.IncrBy(ctx, "id", 1).Result()
	fmt.Println(s, err)
	redis.Z

}
