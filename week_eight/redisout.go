package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

// generateVal 构造不同大小的字节 value
func generateVal(n int) string {
	step := "a"
	var builder strings.Builder
	builder.Grow(n * len(step))
	for i := 0; i < n; i++ {
		builder.WriteString(step)
	}
	return builder.String()
}

func main() {
	redisOptions := &redis.Options{
		Addr:               "127.0.0.1:6379",
		DB:                 0,
		DialTimeout:        time.Second * 10,
		ReadTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		PoolSize:           15,
		PoolTimeout:        time.Second * 30,
		IdleTimeout:        time.Millisecond * 500,
		IdleCheckFrequency: time.Millisecond * 500,
		Password:           "",
	}

	redisClient := redis.NewClient(redisOptions)
	// 写入 value 大小为 100 字节
	val := generateVal(100)
	// 写入数据量为 1w
	dataSize := 10000
	for i := 0; i < dataSize; i++ {
		key := fmt.Sprintf("%d", i)
		rs := redisClient.Set(key, val, time.Second*100)
		err := rs.Err()
		if err != nil {
			fmt.Println("set redis err:", err, key)
			return
		}
	}
}
