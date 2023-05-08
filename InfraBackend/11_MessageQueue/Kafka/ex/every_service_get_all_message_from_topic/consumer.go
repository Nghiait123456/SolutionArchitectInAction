package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type DataRead struct {
	Data1 string `json:"data1"`
	Data2 int    `json:"data2"`
}

func consumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:  "consumer-group-3", // if one service want get all topic from produce, every service must have special GroupID
		Topic:    "my-new-topic-test",
		MinBytes: 10e3, // 10KBl
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Printf("message at topic: %v/partition: %v /offset: %v /key: %s , value  = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		var out DataRead
		errUJ := json.Unmarshal(m.Value, &out)
		if errUJ != nil {
			fmt.Println("Error:", errUJ)
		}

		fmt.Println("data1", out.Data1)
		fmt.Println("data2", out.Data1)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func main() {
	go consumer()
	time.Sleep(1000 * time.Second)
}
