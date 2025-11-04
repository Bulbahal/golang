package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// 3. **Реализуйте функции:**
//   - `SetCache(key string, value string, ttl time.Duration) error` — сохранить значение по ключу с TTL.
//   - `GetCache(key string) (string, error)` — получить значение по ключу.
//   - `DelCache(key string) error` — удалить значение по ключу.
var (
	rd  *redis.Client
	ctx = context.Background()
)

func initRedis() error {
	rd = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	ping, err := rd.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println(ping)
	return nil
}

func SetCache(key, value string, ttl time.Duration) error {
	_, err := rd.Set(ctx, key, value, ttl).Result()
	return err
}

func GetCache(key string) (string, error) {
	val, err := rd.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func DelCache(key string) error {
	return rd.Del(ctx, key).Err()
}

func main() {
	if err := initRedis(); err != nil {
		panic(err)
	}
	err := SetCache("foo", "bar", 10*time.Second)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	val, err := GetCache("foo")
	fmt.Println("foo:", val, err)
	err = DelCache("foo")
	fmt.Println("del foo:", err)
	val, err = GetCache("foo")
	fmt.Println("foo после удаления:", val, err)
}
