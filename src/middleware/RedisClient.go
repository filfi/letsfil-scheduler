package middleware

import (
	"context"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RedisClient struct {
}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (this *RedisClient) RedisClient() *redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     configs.GetValue(configs.RedisHost, configs.Get().Redis.Host) + ":" + configs.GetValue(configs.RedisPort, configs.Get().Redis.Port),
		Password: configs.GetValue(configs.RedisPassword, configs.Get().Redis.Password), // 密码
		DB:       configs.Get().Redis.Database,                                          // 默认数据库
		PoolSize: configs.Get().Redis.PoolSize,                                          // 连接池大小
	})
	// 设置连接超时时间
	fmt.Printf("%v+\n", configs.Get())
	duration, err := time.ParseDuration(configs.Get().Redis.Timeout)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), duration)
	defer cancelFunc()
	result, err := rc.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("redis connect error....")
	}
	fmt.Println("redis: " + result)
	return rc
}
