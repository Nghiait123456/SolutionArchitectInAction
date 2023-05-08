package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

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
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte("test"),
			Value: []byte("one!"),
		},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
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
