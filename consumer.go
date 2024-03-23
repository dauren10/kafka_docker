package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Define Kafka broker address
	kafkaBroker := "localhost:9093"

	// Create a new Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{kafkaBroker},
		Topic:       "your-topic2",
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		MaxWait:     250 * time.Millisecond,
		StartOffset: kafka.FirstOffset,
	})
	defer reader.Close()

	// Trap SIGINT and SIGTERM to trigger a graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Consume messages
	for {
		select {
		case <-signals:
			log.Println("Shutting down")
			return
		default:
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				if err == context.DeadlineExceeded {
					// Timeout error, continue to receive messages
					continue
				}
				log.Printf("Error reading message: %s\n", err)
				return
			}

			log.Printf("Received message: %s\n", string(msg.Value))
		}
	}
}
