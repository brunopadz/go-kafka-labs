package main

import (
	"context"
	"fmt"
	"os"
	"time"

	guuid "github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := os.Getenv("KAFKA_TOPIC")
	boostrap_servers := os.Getenv("KAFKA_BROKER")

	writer := getWriter(boostrap_servers, topic)
	defer writer.Close()

	for {
		// fmt.Println(topic)
		// fmt.Println(boostrap_servers)

		msg := kafka.Message{
			Key:   []byte(guuid.New().String()),
			Value: []byte(guuid.New().String()),
		}

		err := writer.WriteMessages(context.Background(), msg)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second)
	}
}

func getWriter(bootstrap_servers, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{bootstrap_servers},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}
