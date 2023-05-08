package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

// Tạo ra một struct để lưu trữ dữ liệu JSON của bạn.
type DataSend struct {
	Data1 string `json:"data1"`
	Data2 int    `json:"data2"`
}

func producer() {
	// to produce messages
	topic := "my-new-topic-test"
	partition := 0

	fmt.Println("in producer")
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	fmt.Println("start send message")
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	dataS := DataSend{
		Data1: "data1",
		Data2: 10,
	}
	// Chuyển đổi đối tượng JSON sang mảng byte.
	jsonData, err := json.Marshal(dataS)
	if err != nil {
		panic(err)
	}

	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte("one"),
			Value: jsonData,
		},
		kafka.Message{
			Key:   []byte("true"),
			Value: jsonData,
		},
		kafka.Message{
			Key:   []byte("three"),
			Value: jsonData,
		},
	)

	fmt.Println("sen message done")
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func main() {
	producer()
}
