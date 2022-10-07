package main

import (
	"context"
	"encoding/json"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"log"
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
	key1 := "sampleKey"
	value1 := &valueEx{Name: "someName", Email: "someemail@abc.com"}
	err := redisCluterClient.setKey(key1, value1, time.Minute*1)
	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
		panic("error set Key")
	}

	value2 := &valueEx{}
	key2 := "key"
	err2 := redisCluterClient.getKey(key2, value2)
	if err2 != nil {
		fmt.Printf("Error: %v \n", err2.Error())
		panic("error set Key")
	}
	log.Printf("Name: %s", value2.Name)
	log.Printf("Email: %s", value2.Email)
}

func main() {
	ExampleClient()
}
