package main

import (
	"context"
	"encoding/json"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var ctx = context.Background()
var (
	client = &redisCluterClient{}
)

//RedisClusterClient struct
type redisCluterClient struct {
	c *redis.ClusterClient
}

//GetClient get the redis client
func initialize(hostnames string) *redisCluterClient {
	addr := strings.Split(hostnames, ",")
	addr = []string{"testredis.odfdzo.clustercfg.apse1.cache.amazonaws.com:6379"}
	c := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addr,
	})
	if err := c.Ping(ctx).Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	client.c = c
	return client
}

//GetKey get key
func (client *redisCluterClient) getKey(key string, src interface{}) error {
	val, err := client.c.Get(ctx, key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
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
	//Use your actually ip address here
	redisCluterClient := initialize("testredis.odfdzo.clustercfg.apse1.cache.amazonaws.com:6379")
	//key1 :=
	//value1 := &valueEx{Name: "someName", Email: "someemail@abc.com"}

	err := redisCluterClient.c.Set(ctx, "aaa", "tessst", time.Minute*1).Err()
	if err != nil {
		switch {
		case err == redis.Nil:
			fmt.Println("key does not exist")
		case err != nil:
			fmt.Println("Get failed", err)
			panic("exist")
		}
	}

	err2 := redisCluterClient.c.Set(ctx, "key", "ddddd", time.Minute*1).Err()
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

	fmt.Println("value aaa=", redisCluterClient.c.Get(ctx, "aaa"))
	fmt.Println("value key=", redisCluterClient.c.Get(ctx, "key"))

	fmt.Println("test success")
}

func main() {
	ExampleClient()
}
