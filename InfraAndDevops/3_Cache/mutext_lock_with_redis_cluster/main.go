package main

import (
	"context"
	"fmt"
	"github.com/Nghiait123456/redlock"
	"github.com/Nghiait123456/redlock/redis/goredis/v8"
	goredislib "github.com/go-redis/redis/v8"
	"time"
)

func main() {
	//server, err := tempredis.Start(tempredis.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//defer server.Term()

	//client := goredislib.NewClient(&goredislib.Options{
	//	Network: "unix",
	//	Addr:    server.Socket(),
	//})

	client := goredislib.NewClusterClient(&goredislib.ClusterOptions{
		Addrs:    []string{"127.0.0.1:6379"},
		Password: "bitnami",
	})

	pool := goredis.NewPool(client)

	rs := redsync.New(pool)

	mutex := rs.NewMutex("test-redsync")
	ctx := context.Background()

	fmt.Println("start lock")
	if err := mutex.LockContext(ctx); err != nil {
		fmt.Println("lock fail")
		panic(err)
	}
	fmt.Println("start lock success")

	fmt.Println("start race condition lock 1st")
	go func() {
		fmt.Println("start race conditions lock 1st")
		if err := mutex.LockContext(ctx); err != nil {
			fmt.Printf("race conditions fail 1st, err: %v \n", err.Error())
		}
		fmt.Println("race conditions lock success 1st")
	}()

	time.Sleep(10 * time.Second)

	fmt.Println("start end lock")
	if _, err := mutex.UnlockContext(ctx); err != nil {
		fmt.Printf("race conditions fail 1st, err: %v \n", err.Error())
		panic(err)
	}

	fmt.Println("start race condition lock 2st")
	go func() {
		fmt.Println("start race conditions lock 2st")
		if err := mutex.LockContext(ctx); err != nil {
			fmt.Println("race conditions fail 2st")
			panic(err)
		}
		fmt.Println("race conditions lock success 2st")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("end lock success")
}
