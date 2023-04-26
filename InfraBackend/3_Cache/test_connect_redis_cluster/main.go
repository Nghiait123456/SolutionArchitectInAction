package main

import (
	"context"
	"encoding/json"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

//RedisClusterClient struct
type redisCluterClient struct {
	c *redis.ClusterClient
}

type Config struct {
	Addrs    []string
	Password string
}

//GetClient get the redis client
func initialize(cf Config) *redisCluterClient {
	c := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cf.Addrs,
		Password: cf.Password,
	})
	//
	//if err := c.Ping(ctx).Err(); err != nil {
	//	panic("Unable to connect to redis " + err.Error())
	//}

	client := redisCluterClient{
		c: c,
	}
	client.c = c
	return &client
}

//GetKey get key
func (client *redisCluterClient) getKey(key string, value interface{}) error {
	val, err := client.c.Get(ctx, key).Result()
	if err == redis.Nil || err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), &value)
	if err != nil {
		return err
	}

	fmt.Println("000000000000", val, value)
	return nil
}

//SetKey set key
func (client *redisCluterClient) setKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(ctx, key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

type valueEx struct {
	Name  string
	Email string
}

func ExampleClient() {
	redisCluterClient := initialize(Config{
		Addrs:    []string{"127.0.0.1:6379"},
		Password: "bitnami",
	})

	err0 := redisCluterClient.c.Set(ctx, "ValueEmpty", "", time.Minute*1).Err()
	if err0 != nil {
		switch {
		case err0 == redis.Nil:
			fmt.Println("key does not exist")
		case err0 != nil:
			fmt.Println(err0.Error())
			panic("exist")
		}
	}

	err1 := redisCluterClient.c.Set(ctx, "aaa", "tessst", time.Minute*1).Err()
	if err1 != nil {
		switch {
		case err1 == redis.Nil:
			fmt.Println("key does not exist")
		case err1 != nil:
			fmt.Println(err1)
			panic("exist")
		}
	}

	err2 := redisCluterClient.c.Set(ctx, "key", "ddddd", time.Second*1).Err()
	if err2 != nil {
		switch {
		case err2 == redis.Nil:
			fmt.Println("key does not exist")
		case err2 != nil:
			fmt.Println("Get failed", err2.Error())
			panic("exist")
		}

		panic("exist")
	}

	fmt.Println("auth Set Object")

	err3 := redisCluterClient.setKey("keyString", &valueEx{
		Email: "email",
		Name:  "name",
	}, time.Minute*1)
	if err3 != nil {
		switch {
		case err3 == redis.Nil:
			fmt.Println("key does not exist")
		case err3 != nil:
			fmt.Println("Get failed", err3.Error())
			panic("exist")
		}

		panic("exist")
	}

	var value3 valueEx
	err4 := redisCluterClient.getKey("keyString", &value3)
	if err4 != nil {
		switch {
		case err4 == redis.Nil:
			fmt.Println("key does not exist")
		case err4 != nil:
			fmt.Println("Get failed", err4.Error())
			panic("exist")
		}

		panic("exist")
	}

	fmt.Println("get key keyString success ===", value3)

	fmt.Println("value aaa=", redisCluterClient.c.Get(ctx, "aaa"))
	fmt.Println("value key=", redisCluterClient.c.Get(ctx, "key"))
	notExistValue, notExistErr := redisCluterClient.c.Get(ctx, "keyasdasdad").Result()
	fmt.Println("notExistValue", notExistValue, "not existError", notExistErr.Error())

	fmt.Println("start delete varible in cache map key aaa")
	redisCluterClient.c.Del(ctx, "aaa")
	fmt.Println("value aaa=", redisCluterClient.c.Get(ctx, "aaa"))

	fmt.Println("start tets invalidate in cache map key keys")
	time.Sleep(5 * time.Second)
	fmt.Println("value key=", redisCluterClient.c.Get(ctx, "key"))

	fmt.Println("start auth value Empty")
	fmt.Println("value map key ValueEmpty=", redisCluterClient.c.Get(ctx, "ValueEmpty"))

	fmt.Println("auth success")

}

func main() {
	ExampleClient()
}
